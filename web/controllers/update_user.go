package controllers

import (
	"fmt"
	"net/http"
)

func (app *Application) UpdateUserHandler() func(http.ResponseWriter, *http.Request) {

	return app.hasSavedToken(func(w http.ResponseWriter, r *http.Request, token string) {

		fmt.Println("UpdateUserHandler")

		if r.FormValue("updateUserSubmitted") == "true" {

			userId := r.FormValue("userId")
			userType := r.FormValue("userType")
			name := r.FormValue("name")
			email := r.FormValue("email")
			company := r.FormValue("company")
			occupation := r.FormValue("occupation")
			salary := r.FormValue("salary")

			fmt.Println("UpdateUser == " + userId + " -- " + userType + " -- " + name + " -- " + email + " -- " + company + " -- " + occupation + " -- " + salary)

			fabricUser, err := app.Fabric.SessionUser()

			if err != nil {
				fmt.Println("Error Session User  " + err.Error())
			} else {

				err = fabricUser.UpdateUserData(userId, name, email, company, occupation, salary, userType)

				if err != nil {

					fmt.Println("Error Update User Data = " + err.Error())

				} else {

					fmt.Println("Success Update User Data ")
					http.Redirect(w, r, "/index.html", 302)

				}
			}

		}
	})
}
