// @APIVersion 1.0.0
// @Title Funding Api
// @Description Funding
// @Contact
// @TermsOfServiceUrl
package routers

import (
	"funding/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/product",
			beego.NSInclude(
				&controllers.ProductController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserControllers{},
			),
		),
		beego.NSNamespace("/user/address",
			beego.NSInclude(
				&controllers.AddressController{},
			),
		),
		beego.NSNamespace("/user/cart",
			beego.NSInclude(
				&controllers.CartController{},
			),
		),
		beego.NSNamespace("/user/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
