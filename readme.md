## BeeGoLearn

#### Beego.Controller

 Controller 用 Get 获取 Url参数的时候，key前面记得加上冒号

#### Session

使用 Redis 做 Beego Session 的引擎需要在 main.go导入以下两个包

```go
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/gomodule/redigo/redis"

```

