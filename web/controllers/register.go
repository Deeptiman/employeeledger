package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) RegisterHandler(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println("RegisterHandler : " + r.FormValue("signupSubmitted"))

	if r.FormValue("signupSubmitted") == "true" {

		fmt.Println("Register : signupSubmitted")

		nameValue := r.FormValue("name")
		emailValue := r.FormValue("email")
		companyValue := r.FormValue("company")
		occupationValue := r.FormValue("occupation")
		salaryValue := r.FormValue("salary")

		passwordValue := hash(r.FormValue("password"))

		userType := r.FormValue("usermode")

		fmt.Println("Sign up --->  nameValue = " + nameValue + " , emailValue = " + emailValue + " , companyValue = " + companyValue + " , occupationValue = " + occupationValue + " , salaryValue = " + salaryValue + " , userType = " + userType)

		fabricUser, err := app.Fabric.RegisterUser(nameValue, emailValue, companyValue, occupationValue, salaryValue, passwordValue, userType)

		if err != nil {

			fmt.Println("Web Error ----->>> Unable to Register Error Msg : %v", err)

			fmt.Errorf("Web ----- Unable to Register Error Msg : %s", err.Error())

			data.Error = true
			data.ErrorMsg = err.Error()

		} else {

			data.Username = fabricUser.Username
			data.Success = true

			app.processAuthentication(w, emailValue)

			http.Redirect(w, r, "/index.html", 302)

		}

		data.Response = true
	}

	renderTemplate(w, r, "register.html", data)
}
