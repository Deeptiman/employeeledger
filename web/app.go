package web

import (
	"fmt"
	"net/http"
	"employeeledger/web/controllers"
)

func Serve(app *controllers.Application) {

	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/token_auth", app.TokenAuthHandler())

	http.HandleFunc("/login.html", app.LoginHandler)
	http.HandleFunc("/register.html", app.RegisterHandler)
	http.HandleFunc("/index.html", app.IndexPageHandler())
	http.HandleFunc("/update_user", app.UpdateUserHandler())
	http.HandleFunc("/delete_user", app.DeleteUserHandler())

	http.HandleFunc("/change_password.html", app.OpenChangePwdHandler())
	http.HandleFunc("/change_pwd", app.ChangePwdHandler())

	http.HandleFunc("/logout", app.LogoutHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, "/login.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:6000) ...")
	http.ListenAndServe(":6000", nil)

}
