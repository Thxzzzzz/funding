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

func TestGetProductsByPageAndType(t *testing.T) {
	results, err := models.GetProductsByPageAndType(1, 5, 0)
	if err != nil || len(results) < 1 {
		t.Error()
	}
}
