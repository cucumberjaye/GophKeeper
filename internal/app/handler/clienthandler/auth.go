package clienthandler

import (
	"context"
	"fmt"
	"os"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/cucumberjaye/GophKeeper/internal/pkg/tokens"
	"github.com/go-playground/validator/v10"
)

func (c *KeeperClient) registration(signal chan os.Signal) error {
	var login, password string

	for {
		fmt.Println("Registration\nEnter your login (alphanumeric, len > 3) (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case login = <-c.rch:
		}

		fmt.Println("Enter your password (len > 6) (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case password = <-c.rch:
		}

		err := validator.New().Struct(&models.LoginPasswordValidate{
			Login:    login,
			Password: password,
		})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		reqData := &pb.RegistrationRequest{
			Login:    login,
			Password: password,
		}

		resp, err := c.authClient.Registration(context.Background(), reqData)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(resp.Comment)
		break
	}
	return nil
}

func (c *KeeperClient) authentication(signal chan os.Signal) error {
	var login, password string

	for {
		fmt.Println("Authentication\nEnter your login (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case login = <-c.rch:
		}

		fmt.Println("Enter your password (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case password = <-c.rch:
		}

		reqData := &pb.AuthenticationRequest{
			Login:    login,
			Password: password,
		}

		resp, err := c.authClient.Authentication(context.Background(), reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		c.authToken = resp.Token
		c.userID, _ = tokens.ParseToken(resp.Token)
		break
	}
	return nil
}
