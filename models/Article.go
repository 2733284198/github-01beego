package models

/*
	要加入xml的反射，才能解析出 xml的数据
*/
type Article struct {
	Title string `form:"title" xml:"title"`
	//Title   string `form:"title" `
	Content string `form:"content" xml:"content"`
	//Content string `form:"content" `
}
