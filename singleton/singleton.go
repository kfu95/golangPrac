package singleton

import (
	_ "github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Print(){
	fmt.Printf("lol")
}

var conn *db.DB.Connection

// SetDB sets the DB Singleton
func SetDB(val *db.Connection) {
	conn = val
}

// DB returns the DB Singleton
func DB() *db.Connection {
	return conn
}