package resultModels

import (
	"funding/enums"
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
const ProductContentField string = "id,name,product_type,big_img,small_img," +
	"current_price,target_price,backers,end_time," +
	"verify_status,verify_message,created_at,updated_at"

type ProductContent struct {
	ID            uint64              `json:"product_id"`
	Name          string              `json:"name"`           //产品名
	ProductType   int                 `json:"product_type"`   //产品类型
	BigImg        string              `json:"big_img"`        //大图
	SmallImg      string              `json:"small_img"`      //列表小图
	CurrentPrice  float64             `json:"current_price"`  //当前筹集金额
	TargetPrice   float64             `json:"target_price"`   //目标筹集金额
	Backers       int                 `json:"backers"`        //支持人数
	EndTime       time.Time           `json:"end_time"`       //截止时间
	VerifyStatus  int                 `json:"verify_status"`  //审核状态(1：已通过 2：待审核 3:待提交（默认） 4:未通过 )
	VerifyMessage string              `json:"verify_message"` //审核消息（审核失败的原因）
	FundingStatus enums.FundingStatus `json:"funding_status"` //众筹状态
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

// 产品列表
type ProductList struct {
	PageInfo
	ProductContents []*ProductContent `json:"product_contents"`
}

// 结算所需的产品信息
type CheckoutPkgInfo struct {
	ProductId        uint64    `json:"product_id"`         // 产品ID
	SellerID         uint64    `json:"seller_id"`          // 卖家 ID
	SellerNickname   string    `json:"seller_nickname"`    // 卖家昵称
	ProductPackageId uint64    `json:"product_package_id"` // 套餐ID
	ProductName      string    `json:"product_name"`       // 产品名称
	Price            float64   `json:"price"`              // 单价
	Stock            int       `json:"stock"`              // 库存
	Description      string    `json:"description"`        // 套餐描述
	ImageUrl         string    `json:"image_url"`          // 套餐图片
	EndTime          time.Time `json:"end_time"`           // 截止时间
}

// 产品统计信息
type ProductCountInfo struct {
	SupportPriceCount float64 `json:"support_price_count"` // 累计支持金额
	MaxSupportPrice   float64 `json:"max_support_price"`   // 最高筹集金额
	BackersCount      uint64  `json:"backers_count"`       // 累计支持人数
	MaxBackers        uint64  `json:"max_backers"`         // 单项最高支持人数
}

//// 截止时间信息
//type EndTimeItem struct {
//	ProductId uint64    `json:"product_id"` // 产品 Id
//	EndTime   time.Time `json:"end_time"`   //截止时间
//}
