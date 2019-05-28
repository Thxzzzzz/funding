package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["funding/managerControllers:LicenseController"] = append(beego.GlobalControllerRouter["funding/managerControllers:LicenseController"],
        beego.ControllerComments{
            Method: "GetAllLicense",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:LicenseController"] = append(beego.GlobalControllerRouter["funding/managerControllers:LicenseController"],
        beego.ControllerComments{
            Method: "GetByVerifyStatus",
            Router: `/getByVerifyStatus`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:LicenseController"] = append(beego.GlobalControllerRouter["funding/managerControllers:LicenseController"],
        beego.ControllerComments{
            Method: "GetLicenseById",
            Router: `/licenseById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:LicenseController"] = append(beego.GlobalControllerRouter["funding/managerControllers:LicenseController"],
        beego.ControllerComments{
            Method: "UpdateLicense",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "OrderAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "OrderAll",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "ComplaintOrders",
            Router: `/complaintOrders`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "OrderDelete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "OrderRecover",
            Router: `/recover`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerOrderController"],
        beego.ControllerComments{
            Method: "OrderUpdate",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"],
        beego.ControllerComments{
            Method: "GetProductList",
            Router: `/getProductList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"],
        beego.ControllerComments{
            Method: "GetPkgListByProductId",
            Router: `/pkgListByProductId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"],
        beego.ControllerComments{
            Method: "GetProductById",
            Router: `/productById`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerProductController"],
        beego.ControllerComments{
            Method: "UpdateProduct",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "Info",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "GetInfoByRoleId",
            Router: `/infoByRoleId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "NewUser",
            Router: `/newUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"] = append(beego.GlobalControllerRouter["funding/managerControllers:ManagerUserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/updateUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
