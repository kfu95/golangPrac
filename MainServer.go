package main

import (
  
  //"os"
  "net/http"
  "github.com/gorilla/mux"
  common "./common"
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  singleton "./singleton"
)

var router = mux.NewRouter()

func main(){


  db, err := gorm.Open("postgres", "host=localhost port=5432 user=kevinfu dbname=mydb password=12345 sslmode=disable")
  if err != nil {
    panic("failed to connect database")

  }
  singleton.setDB(db)

  

  router.HandleFunc("/",common.MainPageHandler)
  router.HandleFunc("/login", common.LoginPageHandler).Methods("GET")
  router.HandleFunc("/login", common.LoginHandler).Methods("POST")
  router.HandleFunc("/register",common.RegisterPageHandler).Methods("GET")
  router.HandleFunc("/register",common.RegisterHandler).Methods("POST")
  //router.HandleFunc("/logout",common.LogoutHandler).Methods("POST")


  http.Handle("/", router)

  log.Fatal(http.ListenAndServe(":8082",nil))


}
