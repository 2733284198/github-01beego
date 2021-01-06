package beegomodel

//
type Foodtag struct {
	Id          int    `gorm:"primary_key" json:" - "` //
	Name        string `json:"name"`                   //标签名子
	Add_time    int    `json:"add_time"`               //添加时间
	Update_time int    `json:"update_time"`            //更新时间
	Status      uint   `json:"status"`                 //状态  1有效 2无效
	Shop_id     int    `json:"shop_id"`                //店铺ID
}

func (Foodtag) TableName() string {
	return "tb_wm_food_tag"
}
