package clienthandler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/cucumberjaye/GophKeeper/internal/pkg/utils"
	"github.com/cucumberjaye/GophKeeper/pkg/encryption"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

func (c *KeeperClient) setDataHandler(signal chan os.Signal) {
	for {
		var err error
		fmt.Println("Select number:\n1. SetLoginPasswordData\n2. SetTextData\n3. SetBinaryData\n4. SetBankCardData\n5. Back")
		number := <-c.rch
		switch number {
		case "1":
			err = c.setLoginPasswordData(signal)
		case "2":
			err = c.setTextData(signal)
		case "3":
			err = c.setBinaryData(signal)
		case "4":
			err = c.setBankCardData(signal)
		case "5":
			return
		default:
			continue
		}

		if err == nil {
			break
		}
		fmt.Println(err)
	}
}

func (c *KeeperClient) setLoginPasswordData(signal chan os.Signal) error {
	var login, password, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	for {
		fmt.Println("SetLoginPasswordData\nEnter your login (press ctrl+C to back on previous page):")

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

		fmt.Println("Enter description (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case description = <-c.rch:
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_LoginPassword{
				LoginPassword: &pb.LoginPasswordData{
					Login:        login,
					Password:     password,
					Description:  description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.SetData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)

		login, _ = encryption.Encrypt(login)
		password, _ = encryption.Encrypt(password)
		err = c.repo.SetLoginPasswordsData(models.LoginPasswordData{
			Description:  description,
			Login:        login,
			Password:     password,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}

		break
	}

	return nil

}

func (c *KeeperClient) setTextData(signal chan os.Signal) error {
	var data, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	for {
		fmt.Println("SetTextData\nEnter data (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case data = <-c.rch:
		}

		fmt.Println("Enter description (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case description = <-c.rch:
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_Text{
				Text: &pb.TextData{
					Data:         data,
					Description:  description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.SetData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)

		data, _ = encryption.Encrypt(data)

		err = c.repo.SetTextData(models.TextData{
			Description:  description,
			Data:         data,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}

	return nil
}

func (c *KeeperClient) setBinaryData(signal chan os.Signal) error {
	var description string
	var data []byte

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	for {
		fmt.Println("SetBinatyData\nEnter data (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case strData := <-c.rch:
			data = []byte(strData)
		}

		fmt.Println("Enter description (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case description = <-c.rch:
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_BinData{
				BinData: &pb.BinaryData{
					Data:         data,
					Description:  description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.SetData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)

		data, _ = encryption.EncryptBin(data)

		err = c.repo.SetBinaryData(models.BinaryData{
			Description:  description,
			Data:         data,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}

	return nil
}

func (c *KeeperClient) setBankCardData(signal chan os.Signal) error {
	var number, validThru, cvv, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	for {
		fmt.Println("SetBinatyData\nEnter card's number (press ctrl+C to back on previous page):")
		select {
		case <-signal:
			return ErrBack
		case number = <-c.rch:
		}

		fmt.Println("Enter card's validThru (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case validThru = <-c.rch:
		}

		fmt.Println("Enter card's cvv (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case cvv = <-c.rch:
		}

		fmt.Println("Enter description (press ctrl+C to back on previous page):")

		select {
		case <-signal:
			return ErrBack
		case description = <-c.rch:
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_CardData{
				CardData: &pb.BankCardData{
					Number:       number,
					ValidThru:    validThru,
					Cvv:          cvv,
					Description:  description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.SetData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)

		number, _ = encryption.Encrypt(number)
		validThru, _ = encryption.Encrypt(validThru)
		cvv, _ = encryption.Encrypt(cvv)

		err = c.repo.SetBankCardData(models.BankCardData{
			Description:  description,
			Number:       number,
			ValidThru:    validThru,
			CVV:          cvv,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}

		break
	}
	return nil
}

func (c *KeeperClient) getDataArray(signal chan os.Signal) {
	for {
		dataArray, err := c.repo.GetDataArray(c.userID)

		fmt.Println("Choice data (press ctrl+C to back on previous page):")
		store := make(map[string]any)
		for i := range dataArray {
			switch tp := dataArray[i].(type) {
			case models.LoginPasswordData:
				fmt.Printf("\n%d. Description: %s\n", i+1, tp.Description)
			case models.TextData:
				fmt.Printf("\n%d. Description: %s\n", i+1, tp.Description)
			case models.BinaryData:
				fmt.Printf("\n%d. Description: %s\n", i+1, tp.Description)
			case models.BankCardData:
				fmt.Printf("\n%d. Description: %s\n", i+1, tp.Description)
			}
			store[strconv.Itoa(i+1)] = dataArray[i]
		}

		var number string
		select {
		case <-signal:
			fmt.Println(ErrBack)
			return
		case number = <-c.rch:
		}

		data, ok := store[number]
		if ok {
			fmt.Println("1. GetData\n2. UpdateData\n3. DeleteData\n(press ctrl+C to back on previous page)")
			select {
			case <-signal:
				fmt.Println(ErrBack)
				return
			case number = <-c.rch:
				if number == "1" {
					switch v := data.(type) {
					case models.LoginPasswordData:
						utils.PrintLoginPasswordData(v)
					case models.TextData:
						utils.PrintTextData(v)
					case models.BinaryData:
						utils.PrintBinaryData(v)
					case models.BankCardData:
						utils.PrintBankCardData(v)
					}
				} else if number == "2" {
					switch v := data.(type) {
					case models.LoginPasswordData:
						err = c.updateLoginPasswordData(signal, v)
					case models.TextData:
						err = c.updateTextData(signal, v)
					case models.BinaryData:
						err = c.updateBinaryData(signal, v)
					case models.BankCardData:
						err = c.updateBankCardData(signal, v)
					}
				} else if number == "3" {
					c.deleteData(data)
				} else {
					continue
				}
				if errors.Is(err, ErrBack) {
					continue
				}
			}
		}
	}
}

func (c *KeeperClient) updateLoginPasswordData(signal chan os.Signal, data models.LoginPasswordData) error {
	var login, password, encLogin, encPassword string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	oldLogin, _ := encryption.Decode(data.Login)
	oldPassword, _ := encryption.Decode(data.Password)
	for {

		fmt.Printf("UpdateLoginPasswordData (press ctrl+C to back on previous page)\nOld login: %s\nEnter new login: (to skip press enter)\n", oldLogin)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			login = v
			encLogin, _ = encryption.Encrypt(v)
		}

		fmt.Printf("(press ctrl+C to back on previous page)\nOld password: %s\nEnter new password: (to skip press enter)\n", oldPassword)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			password = v
			encPassword, _ = encryption.Encrypt(v)
		}

		if login == "" && password == "" {
			fmt.Println("error: login and password are empty")
			continue
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_LoginPassword{
				LoginPassword: &pb.LoginPasswordData{
					Login:        login,
					Password:     password,
					Description:  data.Description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.UpdateData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)
		err = c.repo.UpdateLoginPasswordsData(models.LoginPasswordData{
			Description:  data.Description,
			Login:        encLogin,
			Password:     encPassword,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}

	return nil
}

func (c *KeeperClient) updateTextData(signal chan os.Signal, data models.TextData) error {
	var textData, encData string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldData, _ := encryption.Decode(data.Data)
	for {
		fmt.Printf("UpdateTextData (press ctrl+C to back on previous page)\nOld data: %s\nEnter new data:\n", oldData)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			textData = v
			encData, _ = encryption.Encrypt(v)
		}

		if textData == "" {
			fmt.Println("error: data is empty")
			continue
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_Text{
				Text: &pb.TextData{
					Data:         textData,
					Description:  data.Description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.UpdateData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)
		err = c.repo.UpdateTextData(models.TextData{
			Description:  data.Description,
			Data:         encData,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}
	return nil
}

func (c *KeeperClient) updateBinaryData(signal chan os.Signal, data models.BinaryData) error {
	var binData string
	var encData []byte

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldData, _ := encryption.DecodeBin(data.Data)
	for {
		fmt.Printf("UpdateBinaryData (press ctrl+C to back on previous page)\nOld data: %v\nEnter new data:\n", oldData)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			binData = v
			encData, _ = encryption.EncryptBin([]byte(v))
		}

		if binData == "" {
			fmt.Println("error: data is empty")
			continue
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_BinData{
				BinData: &pb.BinaryData{
					Data:         []byte(binData),
					Description:  data.Description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.UpdateData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)
		err = c.repo.UpdateBinaryData(models.BinaryData{
			Description:  data.Description,
			Data:         encData,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}
	return nil
}

func (c *KeeperClient) updateBankCardData(signal chan os.Signal, data models.BankCardData) error {
	var number, validThru, cvv, encNumber, encValidThru, encCvv string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldNumber, _ := encryption.Decode(data.Number)
	oldValidThru, _ := encryption.Decode(data.ValidThru)
	oldCVV, _ := encryption.Decode(data.CVV)
	for {
		fmt.Printf("UpdateBankCardData (press ctrl+C to back on previous page)\nOld number: %s\nEnter new number: (to skip press enter):\n", oldNumber)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			number = v
			encNumber, _ = encryption.Encrypt(v)
		}

		fmt.Printf("(press ctrl+C to back on previous page)\nOld validThru: %s\nEnter new validThru: (to skip press enter)\n", oldValidThru)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			validThru = v
			encValidThru, _ = encryption.Encrypt(v)

		}

		fmt.Printf("(press ctrl+C to back on previous page)\nOld cvv: %s\nEnter new cvv: (to skip press enter)\n", oldCVV)

		select {
		case <-signal:
			return ErrBack
		case v := <-c.rch:
			if v == "" {
				break
			}
			cvv = v
			encCvv, _ = encryption.Encrypt(v)
		}

		if number == "" && validThru == "" && cvv == "" {
			fmt.Println("error: number, validThru and cvv are empty")
			continue
		}

		lastModified := time.Now().Unix()

		reqData := &pb.Value{
			Kind: &pb.Value_CardData{
				CardData: &pb.BankCardData{
					Number:       number,
					ValidThru:    validThru,
					Cvv:          cvv,
					Description:  data.Description,
					LastModified: lastModified,
				},
			},
		}

		resp, err := c.storeClient.UpdateData(ctx, reqData)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(resp.Comment)
		err = c.repo.UpdateBankCardData(models.BankCardData{
			Description:  data.Description,
			Number:       encNumber,
			ValidThru:    encValidThru,
			CVV:          encCvv,
			LastModified: lastModified,
		}, c.userID)
		if err != nil {
			fmt.Println(err)
		}
		break
	}
	return nil
}

func (c *KeeperClient) deleteData(data any) {
	var deleteFunc func(key, userID string) error
	var key string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	switch v := data.(type) {
	case models.LoginPasswordData:
		key = v.Description
		deleteFunc = c.repo.DeleteLoginPasswordData
	case models.TextData:
		key = v.Description
		deleteFunc = c.repo.DeleteTextData
	case models.BinaryData:
		key = v.Description
		deleteFunc = c.repo.DeleteBinaryData
	case models.BankCardData:
		key = v.Description
		deleteFunc = c.repo.DeleteBankCardData
	}

	reqData := &pb.Key{Key: key}
	resp, err := c.storeClient.DeleteData(ctx, reqData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Comment)
	err = deleteFunc(key, c.userID)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *KeeperClient) syncer() {
	c.sync()
	ticker := time.NewTicker(5 * time.Minute)

	for {
		<-ticker.C
		c.sync()
	}
}

func (c *KeeperClient) sync() {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	lastSync, err := c.repo.GetLastSync(c.userID)
	if err != nil {
		log.Debug().Err(err).Send()
	}
	resp, err := c.storeClient.Sync(ctx, &pb.Empty{})
	if err != nil {
		log.Debug().Err(err).Send()
	}

	if resp == nil {
		return
	}

	exsits, err := c.repo.GetAllUserKeys(c.userID)
	if err != nil {
		log.Debug().Err(err).Send()
		return
	}

	for i := range resp.Values {
		switch t := resp.Values[i].Kind.(type) {
		case *pb.Value_LoginPassword:
			if t.LoginPassword.LastModified > lastSync {
				err = c.repo.SetLoginPasswordsData(models.LoginPasswordData{
					Description:  t.LoginPassword.Description,
					Login:        t.LoginPassword.Login,
					Password:     t.LoginPassword.Password,
					LastModified: t.LoginPassword.LastModified,
				}, c.userID)
			}
			delete(exsits, t.LoginPassword.Description)
		case *pb.Value_Text:
			if t.Text.LastModified > lastSync {
				err = c.repo.SetTextData(models.TextData{
					Description:  t.Text.Description,
					Data:         t.Text.Data,
					LastModified: t.Text.LastModified,
				}, c.userID)
			}
			delete(exsits, t.Text.Description)
		case *pb.Value_BinData:
			if t.BinData.LastModified > lastSync {
				err = c.repo.SetBinaryData(models.BinaryData{
					Description:  t.BinData.Description,
					Data:         t.BinData.Data,
					LastModified: t.BinData.LastModified,
				}, c.userID)
			}
			delete(exsits, t.BinData.Description)
		case *pb.Value_CardData:
			if t.CardData.LastModified > lastSync {
				err = c.repo.SetBankCardData(models.BankCardData{
					Description:  t.CardData.Description,
					Number:       t.CardData.Number,
					ValidThru:    t.CardData.ValidThru,
					CVV:          t.CardData.Cvv,
					LastModified: t.CardData.LastModified,
				}, c.userID)
			}
			delete(exsits, t.CardData.Description)
		}

		if err != nil {
			log.Debug().Err(err).Send()
		}
	}

	for key, delFun := range exsits {
		err = delFun(key, c.userID)
		if err != nil {
			log.Debug().Err(err).Send()
		}
	}

	err = c.repo.SetLastSync(c.userID)
	if err != nil {
		log.Debug().Err(err).Send()
	}

}
