package singleton

import (
	"fmt"

	datamodel "../datamodel"
	"github.com/jinzhu/gorm"
)

var connection gorm.DB

//SetDB sets the DB Singleton
func SetDB(val *gorm.DB) {
	connection = *val
}

// DB returns the DB Singleton
func DB() *gorm.DB {
	return &connection
}

//DBCheckWrite will write into db what needes to be done is write into the db.
func DBCheckWrite(entry *datamodel.User) (bool, error) {

	connection.Create(&entry)
	// singleton.NewRecord(entry)
	// db.Create(&entry)
	fmt.Printf("here")
	return false, nil
}
