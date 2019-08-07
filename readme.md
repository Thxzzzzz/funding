# Funding 众筹系统 
  当前仓库为 Go 后端仓库
  前端仓库地址      https://github.com/Thxzzzzz/funding-front
  管理者前端仓库地址 https://github.com/Thxzzzzz/funding-manager

# 具体功能的设计和实现

## 首页

>数据库操作相关代码在 ./models/product.go 的 GetProductsByPageAndType() 函数
>
>请求处理的代码在 ./controllers/product.go 的 GetHome() 函数
>
>请求的 url 为  后端地址/product/home使用 Get 方法请求，不需要传入参数，也不需要检验登录状态

首页相关的信息主要是由`products`表提供数据查询。

- 轮播图

  首页最上方为一个组数量为五个的轮播图，轮播图现在为展示最新的五个众筹项目信息。

- 统计信息（暂未完成）

  轮播图下方参考京东众筹的设计，展示 累计支持金额,单项最高筹集金额，单项最高支持人数等统计信息

- 热门项目 （暂未完成）

  将所有种类中近期（一个月？）支持人数最多的几个项目展示出来

- 各类众筹产品的精选项目

  其中最新的一个项目将会显示一张大图，其他六项则显示六个小卡片

  

## 发起众筹

- 填写产品信息

  新的众筹要先填写产品信息，产品信息保存后（保存了才有 product_id 方便后续信息保存）才能进入下一步

- 编辑产品详情
  
  在文本编辑器中编辑产品详情，支持插入图片~
  
- 编辑套餐信息
  
  一条一条的添加套餐信息，每条添加后都将保存到后端
  
- 提交审核
  
  将订单状态由待提交（3） 更改为 待审核（2）以便审核员审核 （这里数据库 products 表的定义和之前有些不一样了，注意论文要改）
  

由于管理众筹还没做好，所以可以先访问 http://localhost:9999/#/user/newFunding?productId=11135&step=1 查看已经上传待审核的这一条，这么一来还顺便完成了编辑功能


## 订单相关

### 订单列表

未付款但是众筹已经结束的订单会显示“未付款，众筹已结束” 并将付款入口隐藏



### 新增订单 

> 数据库操作相关代码在 ./models/order.go 的 NewOrderFromForm() 函数
>
> 请求处理的代码在 ./controllers/order.go 的 OrderList() 函数
>
> 请求的 url 为  后端地址/order/orderList 使用 Post 方法请求 需要传入相关表单数据，需要校验登录状态
>
> 表单相关定义在 ./forms/order.go

需要注意的是，新增订单的时候地址不能保存地址的 ID ，而应该保存姓名、地址、手机号等完整信息，因为地址在之后是有可能被删除的。

新增订单将会从提交的对应信息中（这里是一个列表）创建对应的订单数据，通过数据库事务一次性提交，提交后返回所有提交的订单信息，以便之后的支付订单状态修改。



由于众筹系统一般是没有购物车的，但是这里前端加入购物车有一个酷炫的动画，不舍得不用，所以在结算的时候后端会根据每一个套餐产生一个订单号（每个套餐对应一个订单），然后将订单号列表返回给前端，付款的时候再根据这个订单号列表即可完成多个订单商品同时查询和批量付款。



同理在查看商品详情页的时候实际上也是通过一个订单号列表来进行查询单个订单（因为没有必要再做一个只查询一个订单的接口，能查多个那肯定能查单个）



### 支付

支付这里还没有接入第三方支付平台（本来是要接入支付宝的，但是时间不够了），所以这里点击支付之后会做几件事：
- 发送订单 ID 给后端，请求后端支付对应订单

- 根据订单 ID 查询出对应订单信息

- 将对应订单状态更新为已支付

- 根据订单查询出对应产品

- 检查产品是否已超过众筹时间（众筹结束）如果已结束则回归之前的修改，并返回错误信息，支付失败，否则进入下一步

- 增加对应产品支持者人数，增加已筹集金额

- 根据订单查询出对应套餐

- 检查套餐库存是否足够，如果不够则回归之前的修改，并返回错误信息，支付失败，否则进入下一步

- 减少对应套餐相应的库存，增加支持人数

- 完成支付，返回 OK
  以上步骤在进行数据库更更改时会在一个数据库事务中进行，如有一步发生错误则将会全部回滚，保证数据正确性

  

## 猜你喜欢（商品推荐）

  ### 设计思路

在每次商品列表的卡片（或者轮播图的大图）被点击打开商品详情页时（不做在详情页打开事件里是因为做在那里的话每次详情页刷新都会累加一次），会将在浏览器的 Local Stroage 里将对应类型的累计点击次数+1，然后每次访问有“猜你喜欢”模块的页面（首页，商品列表，商品详情）的时候，就会从 Local Stroage 里读取出累计点击次数最多的产品类型，然后将类型 `product_type` 和需要查询的数量`num`添加到请求参数中，向后端接口`/product/getProductsRand`发起请求，将返回该类型正在众筹中的随机指定数量的产品信息。

### 前端 Local Storage 的处理（算是一种非常简化的推荐算法？）

> 前端关于 Local Storage 的处理代码在 src\utils\storage.js 里面

这里是将推荐相关的数据保存在 Local Storage 的 `fProductRecommend`字段里面 ，这个字段的值为一个数组，所保存的数据结构如下可以 F12 打开浏览器的开发者工具，在 `Application` 标签下 Stroage 里面的 Local Storage 中的 http://localhost:9999 中查看

```json
[
    {
        "type":1, // 产品类型
        "count":5 // 打开次数
    },
    {
        "type":2,
        "count":6
    }
 ]
```

![1558354100983](.\mdImg\1558354100983.png)

在 `src\utils\storage.js` 的代码中，主要是 `Recommend` 字段相关的几个函数（或者说方法？）来对 `fProductRecommend` 进行操作，

``` javascript
// 对应产品类型点击次数保存在 Local Storage 中的 字段名
const recommendStroeKey = 'fProductRecommend'
/**
 * 给对应类型增加点击数量
 * @param {商品类型} type
 * @param {增加的计数值} count
 */
export const addRecommendCount = (type, count) => {
  // 从浏览器 Local Storage 取出类型推荐计数的数组
  let recommend = getStore(recommendStroeKey)
  // 如果没有就初始化一个空数组
  if (!recommend) {
    recommend = []
  } else {
    // 因为  Local Storage 只能存字符串，所以要将 Json 字符串反序列化成 JavaScript 对象
    recommend = JSON.parse(recommend)
  }
  let typeIndex = -1
  // 遍历查找对应类型
  for (let index = 0; index < recommend.length; index++) {
    const element = recommend[index]
    // 找到之后将对应类型的计数累加
    if (element.type === type) {
      typeIndex = index
      recommend[index].count += count
      break
    }
  }
  // 如果没有找到对应类型，则新建一个对象加入数组中
  if (typeIndex === -1) {
    let newTypeRec = {
      type: type,
      count: count
    }
    Array.push(recommend, newTypeRec)
  }
  // 保存到浏览器 Local Storage
  setStore(recommendStroeKey, recommend)
}

/**
 * 获取点击次数最多的类型
 */
export const getRecommendType = () => {
  // 从浏览器 Local Storage 取出类型推荐计数的数组
  let recommend = getStore(recommendStroeKey)
  // 如果没有就返回 0 后端将从所有类型中推荐
  if (!recommend) return 0
  // 因为  Local Storage 只能存字符串，所以要将 Json 字符串反序列化成 JavaScript 对象
  recommend = JSON.parse(recommend)
  let type = 0
  let maxCount = 0
  // 在数组中遍历找最大值
  for (let index = 0; index < recommend.length; index++) {
    const element = recommend[index]
    if (element.count > maxCount) {
      type = element.type
      maxCount = element.count
    }
  }
  // 返回对应类型
  return type
}

/**
 * 清理推荐计数（注销时清理）
 */
export const clearRecommendStorage = () => {
  removeStore(recommendStroeKey)
}

```

`addRecommendCount` 主要是在` src/components/mallGoods.vue` 里面的`openProduct()`(即打开对应商品的详情页)有调用,还有在 `src/pages/Home/home.vue ` 里面的 `openProduct()` （之前是写成了 `LinkTo()`现在改了）里有调用

`getRecommendType` 就是在 home 、goodsDetails还有 goods 里面 `getRecommendProducts()`（获取推荐产品列表）在向后端发起请求前调用，来决定向后端获取哪个类型的产品

`clearRecommendStorage` 则是在 `src/common/header.vue`里面的 `_loginOut()` 里，在用户注销时会清除这个字段（从 Local Storage中删除）

### SQL 查询语句原型

以获取产品类型为 `1`查询数量为 `4`的查询为例，

``` sql
SELECT * 						   # 获取所有字段
FROM `products` 				   # 从 products 表查询
WHERE (verify_status = 1)		   # 只查询通过验证的产品
AND product_type = 1 			   # 只查询产品类型为 1 的产品
AND end_time > CURRENT_TIMESTAMP() # 只查询截止时间大于当前时间的产品（众筹未结束）
ORDER BY RAND()  				   # 按随机数排序（即可随机查询）
LIMIT 4 						   # 限制返回前四个结果
```

### 后端 Gorm 查询

> 代码在 /models/product.go 里面

这里为了节约数据资源用`resultModels.ProductContent`作为返回值，它与 `models.Product ` 的差别在于前者没有 `detail_html`字段，而后者有，由于这里只是要在列表中显示，而 `detail_html` 数据比较长，所以并不需要它

``` go
// 根据类型随机获取指定数量的产品 ProductContent
func GetProductsRandByTypeAndNum(productType int, num int) ([]resultModels.ProductContent, error) {
	var result []resultModels.ProductContent
	cDb := db
	// 只能查到验证通过的
	cDb = cDb.Where("verify_status = ?", enums.Verify_Success)
	// 只能查到未到期的 截止时间大于当前时间
	cDb = cDb.Where("end_time > CURRENT_TIMESTAMP()")
    // 如果传入的类型值不为零，则查询对应类型，否则不添加这个条件
	if productType > 0 {
		cDb = cDb.Where("product_type = ?", productType)
	}
	// 如果传入的数量值不为零，则限制查询返回结果的数量，否则返回所有符合条件的数据
	if num > 0 {
		cDb = cDb.Limit(num)
	}
    // 只返回指定字段（不包含 detail_html）
    // 表为 products 按随机数排序
    // 将结果赋值到 result 中
	err := cDb.Select(resultModels.ProductContentField).
		Table("products").Order("RAND()").
		Scan(&result).Error
	return result, err
}
```






# BeeGo -- 后端框架

## Controller

### BaseController

在 ./controllers/common.go 下定义了一个 `BaseController` struct
以及实现了一些公共函数。后续其他的 Controller 需要嵌入这个 `BaseController`就可以调用针对于 `BaseController` 的一些公共方法。

比如 `ResponseJson()`

```go
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

### 请求解析

#### Get 请求的字段解析到 Struct

> Parsing query string to struct like forms. · Issue #645 · astaxie/beego
>  https://github.com/astaxie/beego/issues/645

以 `ProductListForm` 的解析为例子

首先是 `ProductListForm` 的定义，这里一定要有 form 的 tag

``` go
// 路径 ./forms/product.go
package forms
// 获取产品列表的表单
type ProductListForm struct {
	Page     int     `form:"page"`      // 页码
	PageSize int     `form:"page_size"` // 每页数量
	Type     int     `form:"type"`      // 产品类型
	Sort     int     `form:"sort"`      // 排序方式
	Status   int     `form:"status"`    // form众筹状态
	PriceGt  float64 `form:"price_gt"`  // 价格大于
	PriceLte float64 `form:"price_lte"` // 价格小于
}
```

在 Controller 下的 写法

``` go
func (c *ProductController) GetProductByPage() {
	// TODO 据页码和其他条件获取产品信息
	form := forms.ProductListForm{}
    // 获取所有 query 数据组成的 map
	values := c.Ctx.Request.URL.Query()
    // 解析到 Struct 中
	if err := beego.ParseForm(values, &form); err != nil {
        ……
	}
}
```



### 遇到的问题

Controller 用 Get 获取 Url 参数的时候，key 前面记得加上冒号

## Session

### Session & Cookie 的作用

由于 Http 协议是一个无状态协议，每次后端收到的请求并不能从连接本身来判断是否为同一个连接。而有些 Api 接口需要判断是否已经登录才可以调用（购买/下单之类），每次请求都带上账号密码显然是不合理的做法，所以需要借助 Session 和 Cookie 来实现状态的保持。

> https://cyc2018.github.io/CS-Notes/#/notes/HTTP?id=cookie Cookie

> https://cyc2018.github.io/CS-Notes/#/notes/HTTP?id=_8-session Session

#### <span id="sessionauth">Session 校验身份的方式设计</span>

这里用到的做法是在登录成功后将 `Session Id` 保存在返回数据的 `Headers` 中的 `Cookies` 里的 `token` 字段之中返回到浏览器，浏览器将 `Session Id` 保存在浏览器本地的 `Cookies` 当中的 `token` 字段里;而后端则将登录成功的用户 ID 赋值到 `Session` 的 Values 里面的 `userId` 字段中并将这个 `Session` 以 `SessionId : Session` 这样 Key-Value 的形式存入 Redis 之中;退出登录时则直接销毁整个 `Session`。

之后的每次请求只要带上了 `Cookies`（通常来说只要浏览器中存有该域名的 Cookies 就会在每次请求时自动加入到 Headers 中） 则后端就能够在需要校验身份的时候取出 `Cookies` 中的 `token` 值，利用这个 `token` 值（即 `SessionId`）从 Redis 中查询出对应的 `userId` 的值，如果存在则说明该连接是经过身份校验的，如果不存在则说明没有通过校验（未登录或身份过期）。

### Session 配置

使用 Redis 做 Beego Session 的引擎需要在 main.go 导入以下两个包

```go
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"

```

> https://beego.me/docs/mvc/controller/session.md BeeGo 官方文档

可以在 ./config/app.conf 里配置相关参数

```ini
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
3. 比较通过 将账号信息写入当前请求对应的 Session 之中（这里用了 Redis 作为引擎,这个 Session 会保存到 Redis 里面，参考 [Session 校验身份的方式的设计](#sessionauth)）
4. 返回 Json 信息

```go
// @Title 登录
// @Description 用账号密码登录
// @Param	username	formData	string		true	"账号"
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
		result = resultModels.ErrorResult(resultModels.FALL, "账号或密码错误")
	}
	//  4.. 返回 Json 信息
	c.ResponseJson(result)
}
```

### 登出（注销）时的 Session 处理

注销则比较简单，直接将对应的 `Session` 销毁（从 Redis 中删除）即可，不用管其是否存在对应的 `userId`,因为只要销毁了，那不管其之前存不存在 `userId`，最终处理的结果都是一样的。

```go
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

### 需要验证身份的接口

| https://beego.me/docs/mvc/controller/controller.md 控制器函数 - beego: 简约 & 强大并存的 Go 应用框架

像购物车，收货地址，订单之类的接口都需要验证身份，如果在每一个函数的开头都写上验证身份的代码就显得比较繁琐，这里可以利用 Beego 的控制器函数的 `Prepare` 函数来实现。

为了方便，先在 `./Controller/common.go` 下定义一个 `VailUserController` 并实现 `Prepare` 函数来进行身份验证

```go
// 所有请求都需要身份验证的 Controller
type VailUserController struct {
	BaseController
	User *models.User
}

// 实现 Prepare 验证身份
func (c *VailUserController) Prepare() {
	userId := c.GetSession(SESSION_USER_KEY)
	var result *models.User
	if userId == nil {
		c.ResponseErrJson(errors.New("没有登录"))
		return
	}
	id, _ := userId.(uint64)
	// 获取当前 Session 中的 userId 字段对应的值
	result, err := models.FindUserById(id)
	if err != nil {
		c.ResponseErrJson(errors.New("没有该用户"))
		return
	}
	fmt.Println(result)
	c.User = result
}
```

之后只要是所有接口都需要验证的 Controller 就可以嵌入这个 `VailUserController` 来完成身份验证

比如 `CartController`

```Go
type CartController struct {
	VailUserController
}
```

## 全局错误处理

| https://beego.me/docs/mvc/controller/errors.md 错误处理 - beego: 简约 & 强大并存的 Go 应用框架

Beego 的默认错误处理是会返回一个 html 页面。像这样
![avatar](https://beego.me/docs/images/401.png)

而这个项目是一个前后端分离的项目，返回一个 HTML 页面显然是不太合理，应该让它返回一个 Json 格式的错误信息。

通过官方文档可以知道，这里可以自定义一个 `ErrorController` 来达到这个目的。

首先在 Controller 目录下创建一个 error.go 文件，内容如下:

```go
package controllers

import (
	"funding/resultModels"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	result := resultModels.ErrorResult(404, "Api Not Found")
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ErrorController) Error501() {
	result := resultModels.ErrorResult(501, "Server Error")
	c.Data["json"] = result
	c.ServeJSON()
}
```

然后要在项目根目录下的 main.go 文件的 main 函数里加上

```go
func main() {
	// 异常处理
	beego.ErrorController(&controllers.ErrorController{})
	……………… 略
}

```

# Gorm -- 数据库 ORM 框架 （该项目里使用 MySQL 引擎）

## Gorm 的软删除

在这个项目中包含有 `delete_at`字段的表将会采用软删除的形式来进行删除，即在调用删除函数时，实际上并没有将该条记录从数据库中删除，而是将这个`delete_at` 字段设置为删除时的时间，之后的查询中只要这个 `delete_at` 不为空，就会在查询时忽略这条数据，这样做的好处比较明显的就是能在数据"删除"之后还可以很方便的对其进行恢复操作。

## Raw() 使用原生 SQL 语句查询

在多表查询且不需要所有返回的字段的时候，使用 Gorm 的函数来构建一个查询会显得麻烦许多，所以在多表查询时采用了原生 SQL 来执行查询，并使用 `Scan()` 函数来将结果保存到对应的数据结构中。

以 `.\models\cart.go` 下的 `GetCartItemByUserIdAndPkgId` 函数为例子，这个函数的作用是从 `carts`、 `products`、 `product_packages` 三个表中查询出前端购物车列表项所需的数据并返回对应的 `CartItem` 结构体。

`CartItem` 定义如下

```go
// 前端所显示的购物车 item
type CartItem struct {
	ID               uint64 `json:"id"`                 // 购物车项ID
	UserId           uint64 `json:"user_id"`            // 用户ID
	ProductPackageId uint64 `json:"product_package_id"` // 套餐ID
	Price            string `json:"price"`              // 单价
	stock            int    `json:"stock"`              // 库存
	Nums             int    `json:"nums"`               // 购买数量
	Checked          bool   `json:"checked"`            // 是否勾选
	ProductId        uint64 `json:"product_id"`         // 产品ID
	ProductName      string `json:"product_name"`       // 产品名称
	Description      string `json:"description"`        // 套餐描述
	ImageUrl         string `json:"image_url"`          // 套餐图片
}


```

因为 `ProductName` 存在 products 表中， `Description` 和 `ImageUrl` 存在 product_packages 表中，而 carts 表只存了其对应的 `ProductPackageId` 所以在查询的时候需要联合 product_packages 和 products 表来进行查询。

需要注意的是因为删除的时候是软删除，所以要检查所有的 `deleted_at` 字段，只返回为空的行。

查询的 SQL 语句是这样的：

```SQL
SELECT
	c.id,c.user_id,c.product_package_id,c.nums,c.checked,pkg.product_id,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description
FROM
	carts c
JOIN
	product_packages pkg ON c.product_package_id = pkg.id
JOIN
	products p ON pkg.product_id = p.id
WHERE
	c.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	c.user_id = ? AND
	c.product_package_id = ?
```

其中 `?` 为占位符，传入的参数会按顺序将 `?` 替换掉

于是定义一个常量 `const sqlGetCartItemByUserIdAndPkgId` 值为上面所述的 SQL 语句，使用 Gorm 的 Row() 函数传入 SQL 语句和参数并执行。

```go
// 返回购物车列表项目
func GetCartItemByUserIdAndPkgId(userId uint64, pkgId uint64) (resultModels.CartItem, error) {
	var result resultModels.CartItem
	// 执行 SQL 语句，并将结果映射到 result 中
	err := db.Raw(sqlGetCartItemByUserIdAndPkgId, userId, pkgId).Scan(&result).Error
	return result, err
}
```
