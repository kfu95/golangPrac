package datamodel

import (
	"fmt"	
	singleton "../singleton"
)
type User struct{
	Name, Email, Password string
}

func DBCheckWrite (entry *User) (bool, error){

	singleton.Print()
	// db.NewRecord(entry) 
	// db.Create(&entry)
	fmt.Printf("here")
	return false , nil
}