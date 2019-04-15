# BeeGo -- 后端框架


## Controller
### BaseController
在 ./controllers/common.go 下定义了一个 `BaseController` struct
 以及实现了一些公共函数。后续其他的 Controller 需要嵌入这个 `BaseController`就可以调用针对于 `BaseController` 的一些公共方法。

 比如 `ResponseJson()`
 ``` go
 package controllers

import (
	"github.com/astaxie/beego"
	"testApi/models"
)

//定义 Controller 基类
type BaseController struct {
	beego.Controller
}

func (c *BaseController) ResponseJson(result models.Result) {
	c.Data["json"] = result
	c.ServeJSON()
}
 ```


### 遇到的问题
 Controller 用 Get 获取 Url参数的时候，key前面记得加上冒号



## Session

### Session & Cookie 的作用
由于 Http 协议是一个无状态协议，每次后端收到的请求并不能从连接本身来判断是否为同一个连接。而有些 Api 接口需要判断是否已经登录才可以调用（购买/下单之类），每次请求都带上账号密码显然是不合理的做法，所以需要借助 Session 和 Cookie 来实现状态的保持。

>https://cyc2018.github.io/CS-Notes/#/notes/HTTP?id=cookie Cookie

>https://cyc2018.github.io/CS-Notes/#/notes/HTTP?id=_8-session Session

#### <span id="sessionauth">Session 校验身份的方式设计</span> 

这里用到的做法是在登录成功后将 `Session Id`  保存在返回数据的 `Headers` 中的 `Cookies` 里的 `token` 字段之中返回到浏览器，浏览器将 `Session Id` 保存在浏览器本地的 `Cookies` 当中的 `token` 字段里;而后端则将登录成功的用户 ID 赋值到 `Session` 的 Values 里面的 `userId` 字段中并将这个 `Session` 以 `SessionId : Session` 这样 Key-Value 的形式存入 Redis 之中;退出登录时则直接销毁整个 `Session`。

之后的每次请求只要带上了 `Cookies`（通常来说只要浏览器中存有该域名的 Cookies 就会在每次请求时自动加入到 Headers 中） 则后端就能够在需要校验身份的时候取出 `Cookies` 中的 `token` 值，利用这个 `token` 值（即 `SessionId`）从 Redis 中查询出对应的 `userId` 的值，如果存在则说明该连接是经过身份校验的，如果不存在则说明没有通过校验（未登录或身份过期）。


### Session 配置
使用 Redis 做 Beego Session 的引擎需要在 main.go导入以下两个包

``` go
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"

```
> https://beego.me/docs/mvc/controller/session.md BeeGo 官方文档

可以在 ./config/app.conf 里配置相关参数
``` ini
#开启 Session
sessionon = true
#使用 redis 作为 Session 引擎
sessionprovider = redis
#Session 引擎对应的地址或保存路径，这里为Redis的地址和端口号
sessionproviderconfig = 127.0.0.1:6379
#设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID
sessionname = "token"
```
### 登录时的 Session 处理

登录的主要流程如下：

1. 从请求的数据中获取账号密码
2. 与数据库中的账号密码进行比较
3. 比较通过 将账号信息写入当前请求对应的 Session 之中（这里用了 Redis 作为引擎,这个 Session 会保存到 Redis 里面，参考   [Session 校验身份的方式的设计](#sessionauth)）
4. 返回 Json 信息

``` go
// @Title 登录
// @Description 用账号密码登录
// @Param	username	formData	string		true	"用户名"
// @Param	password	formData	string		true	"密码"
// @Success 200
// @Failure 400
// @router /login [post]
func (c *UserControllers) Login() {
	// 1. 首先获取请求中的数据
	//先声明一个 struct 其结构对应请求的 form 表单数据
	loginForm := forms.LoginForm{}
	var result resultModels.Result
	//将 RequestBody 的值填充到 struct 之中
	err := c.ParseForm(&loginForm)
	//如果解析时出现错误，则说明请求的参数有误
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	}

	// 2. 获取数据库中的数据并与请求数据进行比较
	dbResult, err := models.FindUserByUsername(loginForm.Username)

	//数据库查找出错则返回错误
	if err != nil {
		result = resultModels.ErrorResult(resultModels.FALL, err.Error())
	}

	// 3. 比较得出结果后，如果正确登录则将信息加入到 Session 中
	if dbResult.Password == loginForm.Password {
		result = resultModels.SuccessResult(nil)
		//向当前 Session 写入 userId
		c.SetSession(SESSION_USER_KEY, dbResult.ID)
		//TODO 单点登录
	} else {
		// 密码不正确也返回错误
		result = resultModels.ErrorResult(resultModels.FALL, "用户名或密码错误")
	}
	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}
```

### 登出（注销）时的 Session 处理
注销则比较简单，直接将对应的 `Session` 销毁（从 Redis 中删除）即可，不用管其是否存在对应的 `userId`,因为只要销毁了，那不管其之前存不存在 `userId`，最终处理的结果都是一样的。

``` go
// @Title 登出
// @router /logout	[post]
func (c *UserControllers) Logout() {
	var result models.Result
	//直接销毁 Session
	c.DestroySession()
	result = models.SuccessResult(nil)
	c.ResponseJson(result)
}
```



# Gorm   -- 数据库 ORM 框架

## Gorm 的软删除

在这个项目中包含有 `delete_at`字段的表将会采用软删除的形式来进行删除，即在调用删除函数时，实际上并没有将该条记录从数据库中删除，而是将这个`delete_at` 字段设置为删除时的时间，之后的查询中只要这个 `delete_at` 不为空，就会在查询时忽略这条数据，这样做的好处比较明显的就是能在数据"删除"之后还可以很方便的对其进行恢复操作。