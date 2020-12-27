package models

import (
//"encoding/json"
//"net/http"
)

/**
https://www.jianshu.com/p/2d1fcfcea602
  1.需求
    编写返回商品列表数据的Api，返回数据对应的 struct (go 中struct为值类型)如下所示
*/

//定义接口返回的核心数据 如商品详情信息
type ProductInfo struct {
	Name         string `json:"name"`         //商品名称
	Des          string `json:"des"`          //商品信息描述
	UseMethod    string `json:"useMethod"`    //使用方法
	UseRange     string `json:"useRange"`     //使用范围
	ValidDate    string `json:"validDate"`    //有效期
	RpbStatement string `json:"rpbStatement"` //权责声明
}

type Data struct {
	ProductsInfo []ProductInfo `json:"productsInfo"`
}

//定义接口返回的data数据
type ApiData struct {
	//［结构体变量名 ｜ 变量类型 ｜ json 数据 对应字段名]
	ErrCode int    `json:"errCode"` //接口响应状态码
	Msg     string `json:"msg"`     //接口响应信息
	Data    Data   `json:"data"`
}

//请求路由(映射)至 productsList() handler
//func productsList(w http.ResponseWriter, r *http.Request) {
func ProductsList() ApiData {
	//1.接口状态码
	errorCode := 0
	//2.接口返回信息
	msg := "get products list success!"

	product1 := ProductInfo{
		"柿饼",
		"来自陕西富平的特产",
		"开箱即食",
		"全国",
		"2018-10-20 至 2018-12-31",
		"与xx公司无关",
	}

	product2 := ProductInfo{
		"金枪鱼",
		"来自深海的美味",
		"红烧、清蒸均可",
		"全国",
		"2018-10-20 至 2018-12-31",
		"与xx公司无关",
	}

	product3 := ProductInfo{
		"鲶鱼",
		"来自淡水的美味",
		"红烧、清蒸均可",
		"全国",
		"2018-10-20 至 2018-12-31",
		"与xx公司无关",
	}

	//3.省略对应 DB(关系、非关系、搜索引擎) 操作，以 hard code 代替
	data := Data{[]ProductInfo{product1, product2, product3}}
	//4.组装接口返回的数据
	apiData := ApiData{errorCode, msg, data}
	//5.接口响应复杂json数据
	//json.NewEncoder(w).Encode(apiData)

	return apiData
}
