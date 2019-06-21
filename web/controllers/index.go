package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/employeeledger/chaincode/model"
)

func (app *Application) IndexPageHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &struct {
			Error        bool
			ErrorMsg     string
			Success      bool
			Response     bool
			Admin        bool
			User         bool
			UserData     *model.UserData
			AllUsersData []model.UserData
			Name         string
			UserType     string
			Username     string
		}{

			Error:        false,
			ErrorMsg:     "",
			Success:      false,
			Response:     false,
			Admin:        false,
			User:         false,
			UserData:     nil,
			AllUsersData: nil,
			Username:     "",
		}

		data.Response = true

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {

			data.Error = true
			data.ErrorMsg = err.Error()

		} else {

			UserData, err := fabricUser.GetUserFromLedger(fabricUser.Username)

			if err != nil {

				data.Error = true
				data.ErrorMsg = err.Error()

			} else {

				fmt.Println(" SessionUser Username - " + UserData.Name)

				fmt.Println(" getUserFromLedger UserType - " + UserData.UserType)

				if strings.TrimRight(UserData.UserType, "\n") == "Admin" {

					allUsersData, err := fabricUser.GetAllUsersFromLedger()
					if err != nil {
						http.Error(w, fmt.Sprintf("Unable to retrieve users from the ledger: %v", err), http.StatusInternalServerError)
						return
					} else {

						data.AllUsersData = allUsersData
						data.Admin = true
					}
				} else {

					data.User = true
				}
				data.Name = UserData.Name
				data.UserData = UserData
				data.Username = fabricUser.Username
				data.Success = true
			}

		}

		renderTemplate(w, r, "index.html", data)

	})
}

