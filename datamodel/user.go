package datamodel

//User account
type User struct {
	Name, Email, Password string
}

//TableName returns table name
func (user User) TableName() string {
	return "accounts"
}