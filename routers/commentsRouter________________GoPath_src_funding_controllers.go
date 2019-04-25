package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "GetAddresses",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "DeleteAddress",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "NewAddress",
            Router: `/new`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:AddressController"] = append(beego.GlobalControllerRouter["funding/controllers:AddressController"],
        beego.ControllerComments{
            Method: "UpdateAddress",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "AddCart",
            Router: `/addCart`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartDel",
            Router: `/cartDel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartEdit",
            Router: `/cartEdit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:CartController"] = append(beego.GlobalControllerRouter["funding/controllers:CartController"],
        beego.ControllerComments{
            Method: "CartList",
            Router: `/cartList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ObjectController"] = append(beego.GlobalControllerRouter["funding/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ObjectController"] = append(beego.GlobalControllerRouter["funding/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ObjectController"] = append(beego.GlobalControllerRouter["funding/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ObjectController"] = append(beego.GlobalControllerRouter["funding/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ObjectController"] = append(beego.GlobalControllerRouter["funding/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:OrderController"] = append(beego.GlobalControllerRouter["funding/controllers:OrderController"],
        beego.ControllerComments{
            Method: "NewOrder",
            Router: `/newOrder`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductByPage",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetProductWithPkg",
            Router: `/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetHome",
            Router: `/home`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetVerifyProduct",
            Router: `/verify`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:ProductController"] = append(beego.GlobalControllerRouter["funding/controllers:ProductController"],
        beego.ControllerComments{
            Method: "VerifyProduct",
            Router: `/verify/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "GetUserById",
            Router: `/id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/controllers:UserControllers"] = append(beego.GlobalControllerRouter["funding/controllers:UserControllers"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
