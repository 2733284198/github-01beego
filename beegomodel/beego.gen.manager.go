package beegomodel

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _ManagerMgr struct {
	*_BaseMgr
}

// ManagerMgr open func
func ManagerMgr(db *gorm.DB) *_ManagerMgr {
	if db == nil {
		panic(fmt.Errorf("ManagerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ManagerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("manager"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ManagerMgr) GetTableName() string {
	return "manager"
}

// Get 获取
func (obj *_ManagerMgr) Get() (result Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ManagerMgr) Gets() (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ManagerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_ManagerMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取
func (obj *_ManagerMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithMobile mobile获取
func (obj *_ManagerMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithEmail email获取
func (obj *_ManagerMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithStatus status获取
func (obj *_ManagerMgr) WithStatus(status bool) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRoleID role_id获取
func (obj *_ManagerMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithAddTime add_time获取
func (obj *_ManagerMgr) WithAddTime(addTime int) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithIsSuper is_super获取
func (obj *_ManagerMgr) WithIsSuper(isSuper bool) Option {
	return optionFunc(func(o *options) { o.query["is_super"] = isSuper })
}

// GetByOption 功能选项模式获取
func (obj *_ManagerMgr) GetByOption(opts ...Option) (result Manager, err error) {
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
func (obj *_ManagerMgr) GetByOptions(opts ...Option) (results []*Manager, err error) {
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
func (obj *_ManagerMgr) GetFromID(id int) (result Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromID(ids []int) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_ManagerMgr) GetFromUsername(username string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromUsername(usernames []string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_ManagerMgr) GetFromPassword(password string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromPassword(passwords []string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("password IN (?)", passwords).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容
func (obj *_ManagerMgr) GetFromMobile(mobile string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromMobile(mobiles []string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_ManagerMgr) GetFromEmail(email string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromEmail(emails []string) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ManagerMgr) GetFromStatus(status bool) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromStatus(statuss []bool) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_ManagerMgr) GetFromRoleID(roleID int) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromRoleID(roleIDs []int) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("role_id IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容
func (obj *_ManagerMgr) GetFromAddTime(addTime int) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("add_time = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromAddTime(addTimes []int) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("add_time IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromIsSuper 通过is_super获取内容
func (obj *_ManagerMgr) GetFromIsSuper(isSuper bool) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_super = ?", isSuper).Find(&results).Error

	return
}

// GetBatchFromIsSuper 批量唯一主键查找
func (obj *_ManagerMgr) GetBatchFromIsSuper(isSupers []bool) (results []*Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_super IN (?)", isSupers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ManagerMgr) FetchByPrimaryKey(id int) (result Manager, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
