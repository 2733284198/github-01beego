package models

/*type User struct {
	Id       int
	Username string `form:"username" json:"username" `
	Password string `form:"password" json:"password" `
	Age      int
}*/

type User struct {
	Id       int
	Username string
	Password string
	Age      int
}

//func (User) TableName() string {
//	return "users"
//}
