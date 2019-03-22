package test

import (
	"github.com/astaxie/beego"
	"path/filepath"
	"runtime"
	"testApi/models"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	models.InitDB()
}

func TestFindUserById(t *testing.T) {
	result, err := models.FindUserById(2)
	if err != nil || result.ID != 2 {
		t.Error()
	}
}

func TestFindUserByUsername(t *testing.T) {
	username := "chixuntech"
	result, err := models.FindUserByUsername(username)
	if err != nil || result.Username != username {
		t.Error()
	}
}
