package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) OpenChangePwdHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &struct {
			Error    bool
			ErrorMsg string
			Success  bool
			Response bool
			Name     string
			Email    string
			UserType string
		}{

			Error:    false,
			ErrorMsg: "",
			Success:  false,
			Response: false,
			Name:     "",
			Email:    "",
			UserType: "",
		}

		data.Response = true

		if r.FormValue("openChangePwdSubmitted") == "true" {

			emailValue := r.FormValue("email")
			nameValue := r.FormValue("name")
			userTypeValue := r.FormValue("usermode")

			data.Name = nameValue
			data.Email = emailValue
			data.UserType = userTypeValue
			data.Success = true

			renderTemplate(w, r, "change_password.html", data)
		}
	})
}

func (app *Application) ChangePwdHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		data := &struct {
			Error    bool
			ErrorMsg string
			Success  bool
			Response bool
			Name     string
			Email    string
			UserType string
		}{

			Error:    false,
			ErrorMsg: "",
			Success:  false,
			Response: false,
			Name:     "",
			Email:    "",
			UserType: "",
		}

		data.Response = true

		fabricUser, err := app.Fabric.SessionUser()

		if err != nil {

			data.Error = true
			data.ErrorMsg = err.Error()

		} else {

			if r.FormValue("changePwdSubmitted") == "true" {

				nameValue := r.FormValue("name")
				emailValue := r.FormValue("email")
				userTypeValue := r.FormValue("userType")

				oldPwdValue := hash(r.FormValue("oldPwd"))
				newPwdValue := hash(r.FormValue("newPwd"))

				err = fabricUser.ChangePassword(emailValue, userTypeValue, oldPwdValue, newPwdValue)

				if err != nil {
					fmt.Println("Error : %s " + err.Error())
					fmt.Errorf("Unable to Change user pwd Error Msg : %s", err.Error())

					data.Error = true
					data.ErrorMsg = err.Error()
					data.Name = nameValue
					data.Email = emailValue
					data.UserType = userTypeValue

					renderTemplate(w, r, "change_password.html", data)

				} else {
					http.Redirect(w, r, "/logout", 302)
				}
			}
		}
	})

}
