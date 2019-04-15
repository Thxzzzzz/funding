package test

import (
	"fmt"
	"funding/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	models.InitDB()
}

func TestGetAddressById(t *testing.T) {
	var addsId uint64 = 1
	ret, err := models.FindAddressById(addsId)
	if err != nil && ret.ID != addsId {
		t.Failed()
	}

	ret, err = models.FindAddressById(addsId + 1)
	if err != nil && ret.ID == addsId {
		t.Failed()
	}

}

func TestGetAddressesByUserId(t *testing.T) {
	var userId uint64 = 20003
	rets, err := models.FindAddressesByUserId(userId)
	if err != nil && rets[0].UserId != userId {
		t.Failed()
	}

	rets, err = models.FindAddressesByUserId(userId + 1)
	if err != nil && rets[0].UserId == userId {
		t.Failed()
	}
}

func TestInsertAddress(t *testing.T) {
	address := models.Address{
		UserId:  20003,
		Name:    "测试大佬2",
		Address: "广西壮族自治区桂林市七星区南方小清华",
		Phone:   "18512345432",
	}
	err := models.InsertAddress(&address)
	if err != nil {
		t.Failed()
	}
}

func TestDeleteAddress(t *testing.T) {
	aId := uint64(2)
	err := models.DeleteAddressById(aId)
	if err != nil {
		t.Failed()
	}
}

func TestUpdateAddress(t *testing.T) {
	aId := uint64(1)
	rand.Seed(time.Now().UnixNano())
	randi := rand.Intn(100)
	name := fmt.Sprintf("李小明%d", randi)

	address := models.Address{
		BaseModel: models.BaseModel{
			ID: aId,
		},
		Name: name,
	}
	err := models.UpdateAddress(&address)
	if err != nil {
		t.Failed()
	}
	ret, err := models.FindAddressById(aId)
	if err != nil {
		t.Failed()
	}
	if ret.Name != name {
		t.Failed()
	}
}
