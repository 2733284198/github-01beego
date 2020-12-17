package models

type User struct {
	Id int
	//Username string `form:"username" json:"username" `
	Username string
	//Password string `form:"password" json:"password" `
	Password string
	Age      int
}

func (User) TableName() string {
	return "users"
}
