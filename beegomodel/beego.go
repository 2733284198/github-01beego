package beegomodel

// Articles [...]
type Articles struct {
	ID      int64  `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	Title   string `gorm:"index:article_index;column:title;type:varchar(100);not null" json:"title"`
	Content string `gorm:"column:content;type:longtext" json:"content"`
}

// Manager [...]
type Manager struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Username string `gorm:"column:username;type:varchar(255)" json:"username"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Mobile   string `gorm:"column:mobile;type:varchar(11)" json:"mobile"`
	Email    string `gorm:"column:email;type:varchar(255)" json:"email"`
	Status   bool   `gorm:"column:status;type:tinyint(1)" json:"status"`
	RoleID   int    `gorm:"column:role_id;type:int(11)" json:"role_id"`
	AddTime  int    `gorm:"column:add_time;type:int(11)" json:"add_time"`
	IsSuper  bool   `gorm:"column:is_super;type:tinyint(1)" json:"is_super"`
}

// Person [...]
type Person struct {
	ID   int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Age  int    `gorm:"column:age;type:int(11)" json:"age"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}

// User2 [...]
type User2 struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Username string `gorm:"primary_key;column:username;type:varchar(255);not null" json:"-"`
	Password string `gorm:"primary_key;column:password;type:varchar(255);not null" json:"-"`
	Age      int    `gorm:"column:age;type:int(11)" json:"age"`
}

// Users [...]
type Users struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Username string `gorm:"primary_key;column:username;type:varchar(255);not null" json:"-"`
	Password string `gorm:"primary_key;column:password;type:varchar(255);not null" json:"-"`
	Age      int    `gorm:"column:age;type:int(11)" json:"age"`
}
