package test

import (
	"funding/forms"
	"funding/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"
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

func TestGetProductsByPageAndType(t *testing.T) {
	results, err := models.GetProductsByPageAndType(1, 5, 0)
	if err != nil || len(results) < 1 {
		t.Error()
	}
}

func TestGetProductList(t *testing.T) {
	form := forms.ProductListForm{
		//Type: 1,
		Name: "手环",
	}
	resutlt, err := models.GetProductList(form)
	if err != nil {
		t.Failed()
	}
	t.Log(resutlt)
}

func TestGetCheckoutPkgInfoFromPkgId(t *testing.T) {
	id := uint64(111111113)
	result, err := models.GetCheckoutPkgInfoFromPkgId(id)
	if err != nil {
		t.Failed()
	}
	t.Log(result)
}

func TestGetProductTypeList(t *testing.T) {
	results, err := models.GetProductTypeList()
	if err != nil {
		t.Fail()
		t.Log(err.Error())
	}
	t.Log(results)
}

func TestGetAllProductCountInfo(t *testing.T) {
	result, err := models.GetAllProductCountInfo()
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	t.Log(result)
}
