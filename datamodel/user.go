package datamodel

//User account
type User struct {
	Name, Email, Password string
}

//return table name
func (user User) TableName() string {
	return "accounts"
}
