// @APIVersion 1.0.0
// @Title Funding Api
// @Description Funding
// @Contact
// @TermsOfServiceUrl
package routers

import (
	"funding/controllers"
	"funding/managerControllers"
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
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/pm", // productManager
			beego.NSInclude(
				&controllers.ProductMangerController{},
			),
		),
		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 管理端 API 路由
	managerNs := beego.NewNamespace("/manager",
		beego.NSNamespace("/license",
			beego.NSInclude(
				&managerControllers.LicenseController{},
			),
		),
		beego.NSNamespace("/product",
			beego.NSInclude(
				&managerControllers.ManagerProductController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&managerControllers.ManagerUserController{},
			),
		),
	)

	beego.AddNamespace(managerNs)

}
