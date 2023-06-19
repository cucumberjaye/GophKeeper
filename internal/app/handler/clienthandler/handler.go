package clienthandler

import (
	"fmt"

	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// KeeperClient - содержит поля необходимые для работы клиента.
type KeeperClient struct {
	authClient  pb.AuthenticationClient
	storeClient pb.StorageClient

	authToken string
	userID    string

	app *tview.Application

	repo ClientRepository
}

// ClientRepository - интерфейс взаимодействия с базой данных.
type ClientRepository interface {
	SetLastSync(userID string) error
	GetLastSync(userID string) (int64, error)
	SetData(data any, userID string) error
	GetDataArray(userID string) ([]any, error)
	GetAllUserKeys(userID string) (map[string]func(string, string) error, error)
	UpdateDataRepository
	DeleteDataRepository
}

// UpdateDataRepository - интерфейс описывающий методы обновления данных в базе.
type UpdateDataRepository interface {
	UpdateLoginPasswordData(data models.LoginPasswordData, userID string) error
	UpdateTextData(data models.TextData, userID string) error
	UpdateBinaryData(data models.BinaryData, userID string) error
	UpdateBankCardData(data models.BankCardData, userID string) error
}

// DeleteDataRepository - интерфейс описывающий методы удаления данных из базы.
type DeleteDataRepository interface {
	DeleteLoginPasswordData(key string, userID string) error
	DeleteTextData(key string, userID string) error
	DeleteBinaryData(key string, userID string) error
	DeleteBankCardData(key string, userID string) error
}

// New - создает объект KeeperClient.
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

// Run - запускает клиента.
func (c *KeeperClient) Run() error {
	c.app = tview.NewApplication()

	go c.syncer()

	if err := c.app.SetRoot(c.createMenuForm(), true).EnableMouse(true).Run(); err != nil {
		log.Error().Err(err).Send()
	}

	return nil
}
