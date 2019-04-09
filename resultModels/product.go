package resultModels

import "funding/models"

type HomeResult struct {
	Name          string         `json:"name"`
	Type          int            `json:"type"`
	SortOrder     int            `json:"sortOrder"`
	Position      int            `json:"position"`
	LimitNum      int            `json:"limitNum"`
	Status        int            `json:"status"`
	Remark        string         `json:"remark"`
	PanelContents []PanelContent `json:"panelContents"`
}

type PanelContent struct {
	models.BaseModel
	PanelID         int    `json:"panelId"`
	Type            int    `json:"type"`
	ProductID       int64  `json:"productId"`
	SortOrder       int    `json:"sortOrder"`
	FullURL         string `json:"fullUrl"`
	PicURL          string `json:"picUrl"`
	SalePrice       int    `json:"salePrice"`
	ProductName     string `json:"productName"`
	SubTitle        string `json:"subTitle"`
	ProductImageBig string `json:"productImageBig"`
}
