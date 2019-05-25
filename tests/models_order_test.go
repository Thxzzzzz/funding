package test

import (
	"fmt"
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
			{ProductID: 11111, ProductPackageID: 111111111, Price: 269, Nums: 2, BuyerId: 20003, SellerID: 20002},
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

func TestCancelOrderByOrderId(t *testing.T) {
	if enums.OrderStatus_Finished < enums.OrderStatus_Paid {
		fmt.Println(">>>>>>")
	}
}

func TestUpdateOrder(t *testing.T) {
	order := models.Order{BaseModel: models.BaseModel{ID: 15}, Name: "李小明谁"}
	err := models.UpdateOrder(&order)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
}
func TestUpdateOrderIncludeDeleted(t *testing.T) {
	order := models.Order{BaseModel: models.BaseModel{ID: 15, DeletedAt: nil}, Name: "李小明谁"}
	err := models.UpdateOrderIncludeDeleted(&order)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
}

func TestRecoverDeletedOrder(t *testing.T) {
	err := models.RecoverDeletedOrder(15)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
}

func TestComplaintOrders(t *testing.T) {
	result, err := models.GetComplaintOrders()
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	t.Log(result)
}
