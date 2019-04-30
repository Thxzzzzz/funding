package resultModels

import (
	"time"
)

//SortOrder     int            `json:"sortOrder"`
//Position      int            `json:"position"`
//Status        int            `json:"status"`
//Remark        string         `json:"remark"`
type HomeResult struct {
	Name            string           `json:"name"`
	Type            int              `json:"type"`
	LimitNum        int              `json:"limit_num"`
	ProductContents []ProductContent `json:"product_contents"`
}

// ProductList 的字段
const ProductContentField string = "id,name,product_type,big_img,small_img,current_price,target_price,backers,end_time"

type ProductContent struct {
	ID           uint64    `json:"product_id"`
	Name         string    `json:"name"`          //产品名
	ProductType  int       `json:"product_type"`  //产品类型
	BigImg       string    `json:"big_img"`       //大图
	SmallImg     string    `json:"small_img"`     //列表小图
	CurrentPrice float64   `json:"current_price"` //当前筹集金额
	TargetPrice  float64   `json:"target_price"`  //目标筹集金额
	Backers      int       `json:"backers"`       //支持人数
	EndTime      time.Time `json:"end_time"`      //截止时间
}

// 产品列表
type ProductList struct {
	PageInfo
	ProductContents []*ProductContent `json:"product_contents"`
}
