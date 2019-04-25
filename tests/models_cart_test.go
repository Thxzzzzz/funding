package test

import (
	"funding/models"
	"funding/resultModels"
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

func TestInsertCart(t *testing.T) {
	cart := models.Cart{
		UserId:           20003,
		ProductPackageId: 111111114,
		Nums:             1,
		Checked:          true,
	}
	err := models.InsertCart(&cart)
	if err != nil {
		t.Failed()
	}
}

func TestFindCartById(t *testing.T) {
	cart, err := models.FindCartById(1)
	if err != nil || cart == nil || cart.ID != 1 {
		t.Failed()
	}
}

func TestFindCartByUserIdAndPkgId(t *testing.T) {
	userId := uint64(20003)
	pkgId := uint64(111111114)

	cart, err := models.FindCartByUserIdAndPkgId(userId, pkgId)
	if err != nil || cart.UserId != userId || cart.ProductPackageId != pkgId {
		t.Failed()
	}
}

func TestGetCartItems(t *testing.T) {
	userId := uint64(20003)
	var results []resultModels.CartItem
	results, err := models.GetCartItems(userId)
	if err != nil || len(results) == 0 {
		t.Failed()
	}
}

func TestGetCartItemByUserIdAndPkgId(t *testing.T) {
	userId := uint64(20003)
	pkgId := uint64(111111114)
	cartItem, err := models.GetCartItemByUserIdAndPkgId(userId, pkgId)
	if err != nil {
		t.Failed()
	}
	t.Log(cartItem)
}

func TestUpdateCartByUserIdAndPkgId(t *testing.T) {
	cart := models.Cart{
		UserId:           uint64(20003),
		ProductPackageId: uint64(111111114),
		Nums:             10,
		Checked:          false,
	}

	err := models.UpdateCartByUserIdAndPkgId(&cart)
	if err != nil {
		t.Failed()
	}
}

func TestDeleteCartByUserIdAndPkgId(t *testing.T) {
	userId := uint64(20003)
	pkgId := uint64(111111114)
	err := models.DeleteCartByUserIdAndPkgId(userId, pkgId)
	if err != nil {
		t.Failed()
	}
}

func TestUpdateAllCheckedStatus(t *testing.T) {
	userId := uint64(20003)
	checked := false
	err := models.UpdateAllCheckedStatus(userId, checked)
	if err != nil {
		t.Failed()
	}
}

func TestDeleteAllCheckedCarts(t *testing.T) {
	userId := uint64(20003)
	err := models.DeleteAllCheckedCarts(userId)
	if err != nil {
		t.Failed()
	}
}
