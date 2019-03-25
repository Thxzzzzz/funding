// @APIVersion 1.0.0
// @Title Funding Api
// @Description Funding
// @Contact
// @TermsOfServiceUrl
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"testApi/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//beego.NSNamespace("/object",
		//	beego.NSInclude(
		//		&controllers.ObjectController{},
		//	),
		//),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserControllers{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
