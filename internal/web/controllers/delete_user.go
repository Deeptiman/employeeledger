package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) DeleteUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		fmt.Println("DeleteUserHandler")

		if r.FormValue("deleteUserSubmitted") == "true" {

			email := r.FormValue("email")

			fmt.Println("DeleteUserHandler : Email = " + email)

			fabricUser, err := app.Fabric.SessionUser()

			if err != nil {
				fmt.Println("Error Remove User " + err.Error())
			} else {

				if len(fabricUser.Username) > 0 {

					err := fabricUser.RemoveUser(email)

					if err != nil {
						fmt.Println("DeleteUserHandler : RemoveUser = Error : " + err.Error())
					} else {

						fmt.Println("Success RemoveUser ")

						http.Redirect(w, r, "/index.html", 302)
					}
				}
			}
		}
	})
}
