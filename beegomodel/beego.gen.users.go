package beegomodel

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_UsersMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithAge age获取
func (obj *_UsersMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_UsersMgr) GetFromID(id int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromID(ids []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_UsersMgr) GetFromUsername(username string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromUsername(usernames []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UsersMgr) GetFromPassword(password string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_UsersMgr) GetFromAge(age int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找
func (obj *_UsersMgr) GetBatchFromAge(ages []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id int, username string, password string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ? AND username = ? AND password = ?", id, username, password).Find(&result).Error

	return
}
