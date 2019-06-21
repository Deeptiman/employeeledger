package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	fabricUser, err := app.Fabric.SessionUser()

	if err != nil {
		fmt.Println("Error Logout " + err.Error())
	} else {

		if len(fabricUser.Username) > 0 {

			delete(savedToken, "token")

			fabricUser.Logout()

			fmt.Println("Success Logout ")

			http.Redirect(w, r, "/login.html", 302)
		}
	}
}
