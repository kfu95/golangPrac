package common

import (
	"fmt"
	"net/http"

	datamodel "../datamodel"
	helpers "../helpers"
	singleton "../singleton"
)

// For initial loading of page

// MainPageHandler  handles loading main page
func MainPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/index.html")
	fmt.Fprintf(response, body)
}

// LoginPageHandler prints out login.html
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/login.html")
	fmt.Fprintf(response, body)
}

// LoginHandler handle the post request for login
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")

	fmt.Printf("username is:  %s \n password is:  %s  \n", name, pass)
	http.Redirect(response, request, "/", 302)
}

// RegisterPageHandler handles the register page load
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/register.html")
	fmt.Fprintf(response, body)
}

// RegisterHandler handles registering a new user
func RegisterHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	email := request.FormValue("email")
	pass := request.FormValue("password")
	passCon := request.FormValue("confirmPassword")

	fmt.Printf("username is: %s \n password is:  %s \n passwordConfirmed is: %s \n email is: %s \n ", name, pass, passCon, email)
	aUser := datamodel.User{name, email, pass}
	works, err := singleton.DBCheckWrite(&aUser)
	fmt.Printf("does it work %t", works)
	if err != nil {
		panic("Writing fails here")
	}

	http.Redirect(response, request, "/", 302)
}
