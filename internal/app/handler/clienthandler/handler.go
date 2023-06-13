package clienthandler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"google.golang.org/grpc"
)

var (
	ErrBack = errors.New("back")
)

type KeeperClient struct {
	authClient  pb.AuthenticationClient
	storeClient pb.StorageClient

	authToken string
	userID    string
	rch       chan string

	repo ClientRepository
}

type ClientRepository interface {
	SetLastSync(userID string) error
	GetLastSync(userID string) (int64, error)
	SetLoginPasswordsData(data models.LoginPasswordData, userID string) error
	SetTextData(data models.TextData, userID string) error
	SetBinaryData(data models.BinaryData, userID string) error
	SetBankCardData(data models.BankCardData, userID string) error
	GetDataArray(userID string) ([]any, error)
	GetAllUserKeys(userID string) (map[string]func(string, string) error, error)
	UpdateLoginPasswordsData(data models.LoginPasswordData, userID string) error
	UpdateTextData(data models.TextData, userID string) error
	UpdateBinaryData(data models.BinaryData, userID string) error
	UpdateBankCardData(data models.BankCardData, userID string) error
	DeleteLoginPasswordData(key string, userID string) error
	DeleteTextData(key string, userID string) error
	DeleteBinaryData(key string, userID string) error
	DeleteBankCardData(key string, userID string) error
}

func New(repo ClientRepository) (*KeeperClient, error) {
	cfg, err := configs.New()
	if err != nil {
		return nil, fmt.Errorf("get configs failed with error: %w", err)
	}

	conn, err := grpc.Dial(cfg.ServerAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("connet to  keeper service failed with error: %w", err)
	}
	authClient := pb.NewAuthenticationClient(conn)
	storeClient := pb.NewStorageClient(conn)

	return &KeeperClient{
		authClient:  authClient,
		storeClient: storeClient,
		repo:        repo,
	}, nil
}

func (c *KeeperClient) reader() {
	r := bufio.NewReader(os.Stdin)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		c.rch <- sc.Text()
	}
}

func (c *KeeperClient) Run() error {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)

	c.rch = make(chan string)
	go c.reader()

	fmt.Println("Application starting!")
	for {
		var err error
		fmt.Println("Select number:\n1. Registaration\n2. Authentication\n3. Exit")
		number := <-c.rch

		switch number {
		case "1":
			if err = c.registration(sigint); err == nil {
				err = c.authentication(sigint)
			}
		case "2":
			err = c.authentication(sigint)
		case "3":
			return nil
		default:
			continue
		}

		if err == nil {
			break
		}
		fmt.Println(err.Error())
	}

	go c.syncer()

	for {
		fmt.Println("Select number:\n1. SetData\n2. GetData\n3. Exit")
		number := <-c.rch

		switch number {
		case "1":
			c.setDataHandler(sigint)
		case "2":
			c.getDataArray(sigint)
		case "3":
			return nil

		default:
			continue
		}
	}
}
