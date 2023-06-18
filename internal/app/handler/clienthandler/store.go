package clienthandler

import (
	"context"
	"fmt"
	"time"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/cucumberjaye/GophKeeper/pkg/encryption"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
)

func (c *KeeperClient) createDataMenuForm() tview.Primitive {
	form := tview.NewTable().SetSelectable(true, false).
		SetCell(0, 0, tview.NewTableCell("Set data").SetAlign(1).SetExpansion(1)).
		SetCell(1, 0, tview.NewTableCell("Get data list").SetAlign(1).SetExpansion(1)).
		SetCell(2, 0, tview.NewTableCell("Exit").SetAlign(1).SetExpansion(1)).
		SetSelectedFunc(func(row, column int) {
			switch row {
			case 0:
				c.app.SetRoot(c.createSetDataForm(), true)
			case 1:
				c.app.SetRoot(c.createGetDataArrayForm(), true)
			case 2:
				c.app.Stop()
			}
		})
	form.SetBorder(true).SetTitle("Data Menu")

	return form
}

func (c *KeeperClient) createSetDataForm() tview.Primitive {
	form := tview.NewTable().SetSelectable(true, false).
		SetCell(0, 0, tview.NewTableCell("Login password data").SetAlign(1).SetExpansion(1)).
		SetCell(1, 0, tview.NewTableCell("Text data").SetAlign(1).SetExpansion(1)).
		SetCell(2, 0, tview.NewTableCell("Binary data").SetAlign(1).SetExpansion(1)).
		SetCell(3, 0, tview.NewTableCell("Bank card data").SetAlign(1).SetExpansion(1)).
		SetCell(4, 0, tview.NewTableCell("Data menu").SetAlign(1).SetExpansion(1)).
		SetSelectedFunc(func(row, column int) {
			switch row {
			case 0:
				c.app.SetRoot(c.createSetLoginPasswordDataForm(), true)
			case 1:
				c.app.SetRoot(c.createSetTextDataForm(), true)
			case 2:
				c.app.SetRoot(c.createSetBinaryDataForm(), true)
			case 3:
				c.app.SetRoot(c.createSetBankCardDataForm(), true)
			case 4:
				c.app.SetRoot(c.createDataMenuForm(), true)
			}
		})
	form.SetBorder(true).SetTitle("Set data menu")

	return form
}

func (c *KeeperClient) createSetLoginPasswordDataForm() tview.Primitive {
	var login, password, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	form := tview.NewForm().
		AddInputField("Description", "", 0, nil, func(text string) {
			description = text
		}).
		AddInputField("Login", "", 0, nil, func(text string) {
			login = text
		}).
		AddInputField("Password", "", 0, nil, func(text string) {
			password = text
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createSetDataForm(), true)
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

			_, err := c.storeClient.SetData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			login, _ = encryption.Encrypt(login)
			password, _ = encryption.Encrypt(password)
			err = c.repo.SetData(models.LoginPasswordData{
				Description:  description,
				Login:        login,
				Password:     password,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			c.app.SetRoot(c.createSetDataForm(), true)
		})

	form.SetBorder(true).SetTitle("Set LoginPassword data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createSetTextDataForm() tview.Primitive {
	var data, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	form := tview.NewForm().
		AddInputField("Description", "", 0, nil, func(text string) {
			description = text
		}).
		AddInputField("Data", "", 0, nil, func(text string) {
			data = text
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createSetDataForm(), true)
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

			_, err := c.storeClient.SetData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			data, _ = encryption.Encrypt(data)
			err = c.repo.SetData(models.TextData{
				Description:  description,
				Data:         data,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			c.app.SetRoot(c.createSetDataForm(), true)
		})

	form.SetBorder(true).SetTitle("Set Text data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createSetBinaryDataForm() tview.Primitive {
	var description string
	var data []byte

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	form := tview.NewForm().
		AddInputField("Description", "", 0, nil, func(text string) {
			description = text
		}).
		AddInputField("Data", "", 0, nil, func(text string) {
			data = []byte(text)
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createSetDataForm(), true)
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

			_, err := c.storeClient.SetData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			data, _ = encryption.EncryptBin(data)
			err = c.repo.SetData(models.BinaryData{
				Description:  description,
				Data:         data,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			c.app.SetRoot(c.createSetDataForm(), true)
		})

	form.SetBorder(true).SetTitle("Set Binary data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createSetBankCardDataForm() tview.Primitive {
	var number, validThru, cvv, description string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	form := tview.NewForm().
		AddInputField("Description", "", 0, nil, func(text string) {
			description = text
		}).
		AddInputField("Number", "", 0, nil, func(text string) {
			number = text
		}).
		AddInputField("ValidThru", "", 0, nil, func(text string) {
			validThru = text
		}).
		AddInputField("CVV", "", 0, nil, func(text string) {
			cvv = text
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createSetDataForm(), true)
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

			_, err := c.storeClient.SetData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			number, _ = encryption.Encrypt(number)
			validThru, _ = encryption.Encrypt(validThru)
			cvv, _ = encryption.Encrypt(cvv)
			err = c.repo.SetData(models.BankCardData{
				Description:  description,
				Number:       number,
				ValidThru:    validThru,
				CVV:          cvv,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			c.app.SetRoot(c.createSetDataForm(), true)
		})

	form.SetBorder(true).SetTitle("Set BankCard data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createGetDataArrayForm() tview.Primitive {
	dataArray, _ := c.repo.GetDataArray(c.userID)

	form := tview.NewTable().SetSelectable(true, false)

	store := make(map[int]any)
	for i := range dataArray {
		switch tp := dataArray[i].(type) {
		case models.LoginPasswordData:
			form.SetCell(i, 0, tview.NewTableCell(tp.Description).SetAlign(1).SetExpansion(1))
		case models.TextData:
			form.SetCell(i, 0, tview.NewTableCell(tp.Description).SetAlign(1).SetExpansion(1))
		case models.BinaryData:
			form.SetCell(i, 0, tview.NewTableCell(tp.Description).SetAlign(1).SetExpansion(1))
		case models.BankCardData:
			form.SetCell(i, 0, tview.NewTableCell(tp.Description).SetAlign(1).SetExpansion(1))
		}
		store[i] = dataArray[i]
	}
	form.SetCell(len(dataArray), 0, tview.NewTableCell("Menu").SetAlign(1).SetExpansion(1))
	form.SetSelectedFunc(func(row, column int) {
		if row == len(dataArray) {
			c.app.SetRoot(c.createDataMenuForm(), true)
		} else {
			switch t := store[row].(type) {
			case models.LoginPasswordData:
				c.app.SetRoot(c.createGetLoginPasswordDataFrom(t), true)
			case models.TextData:
				c.app.SetRoot(c.createGetTextDataFrom(t), true)
			case models.BinaryData:
				c.app.SetRoot(c.createGetBinaryDataFrom(t), true)
			case models.BankCardData:
				c.app.SetRoot(c.createGetBankCardDataFrom(t), true)
			}
		}
	})

	form.SetTitle("Data list")

	return form
}

func (c *KeeperClient) createGetLoginPasswordDataFrom(data models.LoginPasswordData) tview.Primitive {
	login, _ := encryption.Decode(data.Login)
	password, _ := encryption.Decode(data.Password)

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 2, false, false).
		AddTextView("Login", login, 0, 2, false, false).
		AddTextView("Password", password, 0, 2, false, false).
		AddButton("Update", func() {
			c.app.SetRoot(c.createUpdateLoginPasswordDataForm(data), true)
		}).
		AddButton("Delete", func() {
			c.deleteData(data)
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		}).
		AddButton("Back", func() {
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	return form
}

func (c *KeeperClient) createGetTextDataFrom(data models.TextData) tview.Primitive {
	textData, _ := encryption.Decode(string(data.Data))

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddTextView("Data", textData, 0, 0, false, false).
		AddButton("Update", func() {
			c.app.SetRoot(c.createUpdateTextDataForm(data), true)
		}).
		AddButton("Delete", func() {
			c.deleteData(data)
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		}).
		AddButton("Back", func() {
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	return form
}

func (c *KeeperClient) createGetBinaryDataFrom(data models.BinaryData) tview.Primitive {
	binData, _ := encryption.DecodeBin(data.Data)

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddTextView("Data", string(binData), 0, 0, false, false).
		AddButton("Update", func() {
			c.app.SetRoot(c.createUpdateBinaryDataForm(data), true)
		}).
		AddButton("Delete", func() {
			c.deleteData(data)
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		}).
		AddButton("Back", func() {
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	return form
}

func (c *KeeperClient) createGetBankCardDataFrom(data models.BankCardData) tview.Primitive {
	number, _ := encryption.Decode(data.Number)
	validThru, _ := encryption.Decode(data.ValidThru)
	cvv, _ := encryption.Decode(data.CVV)

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddTextView("Number", number, 0, 0, false, false).
		AddTextView("ValidThru", validThru, 0, 0, false, false).
		AddTextView("CVV", cvv, 0, 0, false, false).
		AddButton("Update", func() {
			c.app.SetRoot(c.createUpdateBankCardDataForm(data), true)
		}).
		AddButton("Delete", func() {
			c.deleteData(data)
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		}).
		AddButton("Back", func() {
			c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	return form
}

func (c *KeeperClient) createUpdateLoginPasswordDataForm(data models.LoginPasswordData) tview.Primitive {
	var login, password, encLogin, encPassword string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)

	oldLogin, _ := encryption.Decode(data.Login)
	oldPassword, _ := encryption.Decode(data.Password)

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddInputField("Login", oldLogin, 0, nil, func(text string) {
			login = text
			encLogin, _ = encryption.Encrypt(text)
		}).
		AddInputField("Password", oldPassword, 20, nil, func(text string) {
			password = text
			encPassword, _ = encryption.Encrypt(text)
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
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

			_, err := c.storeClient.UpdateData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = c.repo.UpdateLoginPasswordData(models.LoginPasswordData{
				Description:  data.Description,
				Login:        encLogin,
				Password:     encPassword,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	form.SetBorder(true).SetTitle("Update LooginPassword data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createUpdateTextDataForm(data models.TextData) tview.Primitive {
	var textData, encData string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldData, _ := encryption.Decode(data.Data)
	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddInputField("Data", oldData, 0, nil, func(text string) {
			textData = text
			encData, _ = encryption.Encrypt(text)
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
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

			_, err := c.storeClient.UpdateData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = c.repo.UpdateTextData(models.TextData{
				Description:  data.Description,
				Data:         encData,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	form.SetBorder(true).SetTitle("Set Text data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createUpdateBinaryDataForm(data models.BinaryData) tview.Primitive {
	var binData string
	var encData []byte

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldData, _ := encryption.DecodeBin(data.Data)
	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddInputField("Data", string(oldData), 0, nil, func(text string) {
			binData = text
			encData, _ = encryption.EncryptBin([]byte(text))
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
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

			_, err := c.storeClient.UpdateData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = c.repo.UpdateBinaryData(models.BinaryData{
				Description:  data.Description,
				Data:         encData,
				LastModified: lastModified,
			}, c.userID)
			if err != nil {
				fmt.Println(err)
			}
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	form.SetBorder(true).SetTitle("Set Text data").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createUpdateBankCardDataForm(data models.BankCardData) tview.Primitive {
	var number, validThru, cvv, encNumber, encValidThru, encCvv string

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authentication", c.authToken)
	oldNumber, _ := encryption.Decode(data.Number)
	oldValidThru, _ := encryption.Decode(data.ValidThru)
	oldCVV, _ := encryption.Decode(data.CVV)

	form := tview.NewForm().
		AddTextView("Description", data.Description, 0, 0, false, false).
		AddInputField("Number", oldNumber, 0, nil, func(text string) {
			number = text
			encNumber, _ = encryption.Encrypt(text)
		}).
		AddInputField("ValidThru", oldValidThru, 0, nil, func(text string) {
			validThru = text
			encValidThru, _ = encryption.Encrypt(text)
		}).
		AddInputField("CVV", oldCVV, 0, nil, func(text string) {
			cvv = text
			encCvv, _ = encryption.Encrypt(text)
		}).
		AddButton("OK", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
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

			_, err := c.storeClient.UpdateData(ctx, reqData)
			if err != nil {
				fmt.Println(err)
				return
			}

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
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createGetDataArrayForm(), true)
		})

	form.SetBorder(true).SetTitle("Update LooginPassword data").SetTitleAlign(tview.AlignLeft)

	return form
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
	_, err := c.storeClient.DeleteData(ctx, reqData)
	if err != nil {
		fmt.Println(err)
	}
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
				err = c.repo.SetData(models.LoginPasswordData{
					Description:  t.LoginPassword.Description,
					Login:        t.LoginPassword.Login,
					Password:     t.LoginPassword.Password,
					LastModified: t.LoginPassword.LastModified,
				}, c.userID)
			}
			delete(exsits, t.LoginPassword.Description)
		case *pb.Value_Text:
			if t.Text.LastModified > lastSync {
				err = c.repo.SetData(models.TextData{
					Description:  t.Text.Description,
					Data:         t.Text.Data,
					LastModified: t.Text.LastModified,
				}, c.userID)
			}
			delete(exsits, t.Text.Description)
		case *pb.Value_BinData:
			if t.BinData.LastModified > lastSync {
				err = c.repo.SetData(models.BinaryData{
					Description:  t.BinData.Description,
					Data:         t.BinData.Data,
					LastModified: t.BinData.LastModified,
				}, c.userID)
			}
			delete(exsits, t.BinData.Description)
		case *pb.Value_CardData:
			if t.CardData.LastModified > lastSync {
				err = c.repo.SetData(models.BankCardData{
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
