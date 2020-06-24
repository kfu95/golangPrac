package main

import (
	"log"
	"net/http"
	singleton "./singleton"
	controller "./controller"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var router = mux.NewRouter()

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=kevinfu dbname=mydb password=12345 sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	singleton.SetDB(db)

	router.HandleFunc("/", controller.MainPageHandler)
	router.HandleFunc("/login", controller.LoginPageHandler).Methods("GET")
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	router.HandleFunc("/register", controller.RegisterPageHandler).Methods("GET")
	router.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	//router.HandleFunc("/logout",controller.LogoutHandler).Methods("POST")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8082", nil))

}
