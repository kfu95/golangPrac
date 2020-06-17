package common

import(
	"fmt"
	"net/http"
	helpers "../helpers"
	datamodel "../datamodel"
)



// For initial loading of page

// load of the log in page
func MainPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/index.html")
    fmt.Fprintf(response, body)
}

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/login.html")
    fmt.Fprintf(response, body)
}
// handle the post request 
func LoginHandler(response http.ResponseWriter, request *http.Request){
	name := request.FormValue("name")
	pass := request.FormValue("password")
	
	fmt.Printf("username is:  %s \n password is:  %s  \n", name, pass)
	http.Redirect(response, request ,"/", 302)
}

func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/register.html")
    fmt.Fprintf(response, body)
}

func RegisterHandler(response http.ResponseWriter, request *http.Request){
	name := request.FormValue("name")
	email := request.FormValue("email")
	pass := request.FormValue("password")
	passCon := request.FormValue("confirmPassword")

	fmt.Printf("username is: %s \n password is:  %s \n passwordConfirmed is: %s \n email is: %s \n ", name, pass, passCon, email)
	aUser := datamodel.User{name,email,pass}
	datamodel.DBCheckWrite(&aUser)

	http.Redirect(response, request, "/",302)
}
