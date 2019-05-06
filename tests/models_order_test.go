package test

import (
	"funding/forms"
	"funding/models"
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	models.InitDB()
}

func TestNewOrderFromForm(t *testing.T) {
	form := forms.NewOrderForm{
		Name:       "小明",
		Address:    "桂电",
		Phone:      "18512345678",
		OrderTotal: 538,
		OrderPkgList: []forms.OrderPkgItem{
			{ProductID: 11111, ProductPackageID: 111111111, Price: 269, Nums: 2, UserID: 20003, SellerID: 20002},
		},
	}
	orders, err := models.NewOrderFromForm(20003, &form)
	if err != nil {
		t.Failed()
	}
	t.Log(orders)
}
