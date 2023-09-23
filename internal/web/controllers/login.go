package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {

	data := &struct {
		TransactionId string
		ErrorMsg      string
		Error         bool
		Success       bool
		Response      bool
		Username      string
	}{
		TransactionId: "",
		ErrorMsg:      "",
		Error:         false,
		Success:       false,
		Response:      false,
		Username:      "",
	}

	if r.FormValue("signinSubmitted") == "true" {

		emailValue := r.FormValue("email")
		passwordValue := hash(r.FormValue("password"))

		fmt.Println("Sign In --->  emailValue = " + emailValue)

		fabricUser, err := app.Fabric.LoginUser(emailValue, passwordValue)

		if err != nil {
			fmt.Errorf("Web ----- unable to login user : %v", err)

			//http.Error(w, err.Error(), 500)

			data.Error = true
			data.ErrorMsg = err.Error()
		} else {

			fmt.Println("Logged In User : " + fabricUser.Username)

			fmt.Println("Login : Redirect Index")

			app.processAuthentication(w, emailValue)

			http.Redirect(w, r, "/index.html", 302)

			data.Username = fabricUser.Username
			data.Success = true
		}
		data.Response = true
	}

	renderTemplate(w, r, "login.html", data)
}
