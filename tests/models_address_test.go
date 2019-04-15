package test

import (
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

func TestGetAddressById(t *testing.T) {
	var addsId uint64 = 1
	ret, err := models.GetAddressById(addsId)
	if err != nil && ret.ID != addsId {
		t.Failed()
	}

	ret, err = models.GetAddressById(addsId + 1)
	if err != nil && ret.ID == addsId {
		t.Failed()
	}

}

func TestGetAddressesByUserId(t *testing.T) {
	var userId uint64 = 20003
	rets, err := models.GetAddressesByUserId(userId)
	if err != nil && rets[0].UserId != userId {
		t.Failed()
	}

	rets, err = models.GetAddressesByUserId(userId + 1)
	if err != nil && rets[0].UserId == userId {
		t.Failed()
	}
}
