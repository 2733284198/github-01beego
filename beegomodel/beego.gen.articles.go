package beegomodel

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ArticlesMgr struct {
	*_BaseMgr
}

// ArticlesMgr open func
func ArticlesMgr(db *gorm.DB) *_ArticlesMgr {
	if db == nil {
		panic(fmt.Errorf("ArticlesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArticlesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("articles"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArticlesMgr) GetTableName() string {
	return "articles"
}

// Get 获取
func (obj *_ArticlesMgr) Get() (result Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArticlesMgr) Gets() (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArticlesMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取
func (obj *_ArticlesMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取
func (obj *_ArticlesMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// GetByOption 功能选项模式获取
func (obj *_ArticlesMgr) GetByOption(opts ...Option) (result Articles, err error) {
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
func (obj *_ArticlesMgr) GetByOptions(opts ...Option) (results []*Articles, err error) {
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
func (obj *_ArticlesMgr) GetFromID(id int64) (result Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ArticlesMgr) GetBatchFromID(ids []int64) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_ArticlesMgr) GetFromTitle(title string) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_ArticlesMgr) GetBatchFromTitle(titles []string) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_ArticlesMgr) GetFromContent(content string) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_ArticlesMgr) GetBatchFromContent(contents []string) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ArticlesMgr) FetchByPrimaryKey(id int64) (result Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByArticleIndex  获取多个内容
func (obj *_ArticlesMgr) FetchIndexByArticleIndex(title string) (results []*Articles, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("title = ?", title).Find(&results).Error

	return
}
