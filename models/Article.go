package models

/*
	要加入xml的反射，才能解析出 xml的数据
*/
type Article struct {
	//gorm.Model
	//ID int `gorm:primary_key `
	//ID int

	Id    int
	Title string `form:"title" xml:"title" gorm:"type:varchar(100); not null; index:article_index" `
	//Title   string `form:"title" `
	Content string `form:"content" xml:"content"`
	//Content string `form:"content" `
	Desc string ` gorm:"-" `
}
