package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("thisisahyperledgerfabricdemo")
var savedToken = make(map[string]string)

func (app *Application) processAuthentication(w http.ResponseWriter, key string) {

	validToken, err := GenerateJWT(key)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println("Register : Redirect Index")

	client := &http.Client{}
	r, _ := http.NewRequest("GET", "http://localhost:"+PORT+"/token_auth", nil)
	r.Header.Set("Token", validToken)

	_, err = client.Do(r)

	if err != nil {
		fmt.Fprintf(w, ">>>> Error: %s", err.Error())
	}

	savedToken["token"] = validToken
}

func GenerateJWT(email string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Println("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (app *Application) TokenAuthHandler() func(http.ResponseWriter, *http.Request) {

	return app.isAuthorized(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (app *Application) hasSavedToken(endpoint func(http.ResponseWriter, *http.Request, string)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if savedToken["token"] != "" {
			token := savedToken["token"]
			endpoint(w, r, token)
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	}
}

func (app *Application) isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was a error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	}
}
