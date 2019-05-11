package test

import (
	"funding/enums"
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
		t.Fail()
	}
	t.Log(orders)
}

func TestGetOrderListByOrderIds(t *testing.T) {
	orderIds := []uint64{3, 4, 5}
	userId := uint64(20003)
	roleId := enums.Role_Buyer
	orderItems, err := models.GetOrderListByOrderIds(orderIds, userId, roleId)
	if err != nil {
		t.Fail()
	}
	t.Log(orderItems)
}

func TestPayOrderByOrderIdList(t *testing.T) {
	orderIds := []uint64{8}
	err := models.PayOrderByOrderIdList(orderIds)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestSendOutOrderById(t *testing.T) {
	form := forms.OrderSendOutForm{
		OrderId:        5,
		CheckingNumber: "12345678",
	}
	userId := uint64(20002)
	err := models.SendOutOrderById(&form, userId)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
