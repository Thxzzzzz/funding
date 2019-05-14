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

func TestGetResultCommentInfosByProductId(t *testing.T) {
	form := forms.CommentListByProductForm{ProductId: 11128}
	results, err := models.GetResultCommentInfosByProductId(&form)
	if err != nil {
		t.Fail()
		t.Log(err)
	}
	t.Log(results)

}
