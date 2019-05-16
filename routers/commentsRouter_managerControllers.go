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
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
