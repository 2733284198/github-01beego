package beegomodel

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _PersonMgr struct {
	*_BaseMgr
}

// PersonMgr open func
func PersonMgr(db *gorm.DB) *_PersonMgr {
	if db == nil {
		panic(fmt.Errorf("PersonMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PersonMgr{_BaseMgr: &_BaseMgr{DB: db.Table("person"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PersonMgr) GetTableName() string {
	return "person"
}

// Get 获取
func (obj *_PersonMgr) Get() (result Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PersonMgr) Gets() (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PersonMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAge age获取
func (obj *_PersonMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithName name获取
func (obj *_PersonMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_PersonMgr) GetByOption(opts ...Option) (result Person, err error) {
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
func (obj *_PersonMgr) GetByOptions(opts ...Option) (results []*Person, err error) {
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
func (obj *_PersonMgr) GetFromID(id int) (result Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_PersonMgr) GetBatchFromID(ids []int) (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_PersonMgr) GetFromAge(age int) (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量唯一主键查找
func (obj *_PersonMgr) GetBatchFromAge(ages []int) (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("age IN (?)", ages).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_PersonMgr) GetFromName(name string) (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_PersonMgr) GetBatchFromName(names []string) (results []*Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PersonMgr) FetchByPrimaryKey(id int) (result Person, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
