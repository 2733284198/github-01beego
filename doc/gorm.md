# gorm in操作:

https://www.cnblogs.com/rickiyang/p/11074162.html



	db.Where("name in (?)",[]string{"xiaoming","xiaohong"}).Find(&u)


# 李文周的博客

选择字段
Select，指定你想从数据库中检索出的字段，默认会选择全部字段。

db.Select("name, age").Find(&users)
//// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
//// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
//// SELECT COALESCE(age,'42') FROM users;


