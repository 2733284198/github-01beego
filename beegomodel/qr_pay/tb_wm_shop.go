package qr_pay

//
type Tb_wm_shop struct {
	Id                   int     `gorm:"primary_key" json:" - "` //主键
	Shop_id              int     `json:"shop_id"`                //旧的店铺id(通过这个字段去关联)
	Mobile               string  `json:"mobile"`                 //联系电话
	Password             string  `json:"password"`               //密码
	Shopname             string  `json:"shopname"`               //商品名称
	Login_info           string  `json:"login_info"`             //登录信息
	Num_login_error      uint    `json:"num_login_error"`        //登录错误次数
	Time_login_lock      int     `json:"time_login_lock"`        //锁定登录时间
	Status               uint    `json:"status"`                 //状态
	Add_time             int     `json:"add_time"`               //加入时间
	Address              string  `json:"address"`                //
	Logo                 string  `json:"logo"`                   //店铺logo
	Shop_type            uint    `json:"shop_type"`              //店铺类型：0-总店，1-分店
	Parent_id            int     `json:"parent_id"`              //总店编号
	Email                string  `json:"email"`                  //邮箱
	Send_cost            float64 `json:"send_cost"`              //配送费
	Box_cost             float64 `json:"box_cost"`               //包装费
	Org_shop_id          int     `json:"org_shop_id"`            //原始shop_id
	Reduce_money_config  string  `json:"reduce_money_config"`    //满减的配置
	Is_open_reduce_money uint    `json:"is_open_reduce_money"`   //是否启用满减 0不开启 1开启
	Can_use_range        string  `json:"can_use_range"`          //1  外卖配送 2堂食 多个用逗号分隔开
	Store_template       int     `json:"store_template"`         //店铺模板id
	League_status        uint    `json:"league_status"`          //联盟冻结状态:10=开启,20=关闭
	Activity_img         string  `json:"activity_img"`           //活动图片
	Shop_type_id         int     `json:"shop_type_id"`           //店铺类型ID
	Is_platform_activity uint    `json:"is_platform_activity"`   //是否开启平台活动
	Longitude            string  `json:"longitude"`              //
	Latitude             string  `json:"latitude"`               //
	Deleted_at           int     `json:"deleted_at"`             //
}
