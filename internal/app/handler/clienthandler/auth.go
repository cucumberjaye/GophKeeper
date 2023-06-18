package clienthandler

import (
	"context"
	"fmt"

	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/rivo/tview"
)

func (c *KeeperClient) createMenuForm() tview.Primitive {
	mainForm := tview.NewTable().SetSelectable(true, false).
		SetCell(0, 0, tview.NewTableCell("Registaration").SetAlign(1).SetExpansion(1)).
		SetCell(1, 0, tview.NewTableCell("Authentication").SetAlign(1).SetExpansion(1)).
		SetCell(2, 0, tview.NewTableCell("Exit").SetAlign(1).SetExpansion(1)).
		SetSelectedFunc(func(row, column int) {
			switch row {
			case 0:
				c.app.SetRoot(c.createRegistrationForm(), true)
			case 1:
				c.app.SetRoot(c.createAuthenticationForm(), true)
			case 2:
				c.app.Stop()
			}
		})
	mainForm.SetBorder(true).SetTitle("Menu")

	return mainForm
}

func (c *KeeperClient) createRegistrationForm() tview.Primitive {
	var login, password string

	form := tview.NewForm().
		AddInputField("Login", "", 0, nil, func(text string) {
			login = text
		}).
		AddPasswordField("Password", "", 20, '*', func(text string) {
			password = text
		}).
		AddButton("OK", func() {
			reqData := &pb.RegistrationRequest{
				Login:    login,
				Password: password,
			}
			_, err := c.authClient.Registration(context.Background(), reqData)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			c.app.SetRoot(c.createAuthenticationForm(), true)
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createMenuForm(), true)
		})

	form.SetBorder(true).SetTitle("Registaration").SetTitleAlign(tview.AlignLeft)

	return form
}

func (c *KeeperClient) createAuthenticationForm() tview.Primitive {
	var login, password string

	form := tview.NewForm().
		AddInputField("Login", "", 0, nil, func(text string) {
			login = text
		}).
		AddPasswordField("Password", "", 20, '*', func(text string) {
			password = text
		}).
		AddButton("OK", func() {
			reqData := &pb.AuthenticationRequest{
				Login:    login,
				Password: password,
			}
			resp, err := c.authClient.Authentication(context.Background(), reqData)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			c.authToken = resp.Token

			c.app.SetRoot(c.createDataMenuForm(), true)
		}).
		AddButton("Menu", func() {
			defer c.app.SetRoot(c.createMenuForm(), true)
		})

	form.SetBorder(true).SetTitle("Authentication").SetTitleAlign(tview.AlignLeft)

	return form
}
