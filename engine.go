//	PhalGo-engine
//	注意路由引擎,依赖Echo对器进行封装
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/labstack/echo"

package phalgo

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"net/http"
)

const (
	RETJSON = 1
	RETMXL = 2
	RETDES = 3
)

var (
	Echo *echo.Echo
	RetType int = 1
)
// 初始化echo实例
func NewEcho() *echo.Echo {

	Echo = echo.New()
	return Echo
}

// 设置Ret格式
func SetRetType(i int) {

	RetType = i
}


// 开启服务
func Start(prot string) {
	Echo.Logger.Fatal(Echo.Start(":1323"))
}

// 打印请求异常信息
func Recover() {

	Echo.Use(middleware.Recover())
}

// 是否开启debug
func SetDebug(on bool) {
	Echo.Debug = on
}

// 获取debug状态
func Debug() bool {
	return Echo.Debug
}

// 打印请求信息
func Logger() {

	Echo.Use(middleware.Logger())
}

// 开启gzip压缩
func Gzip() {

	Echo.Use(middleware.Gzip())
}

// 设置Body大小
func BodyLimit(str string) {

	Echo.Use(middleware.BodyLimit(str))
}

// 自动添加末尾斜杠
func AddTrailingSlash() {

	Echo.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}

// 自动删除末尾斜杠
func RemoveTrailingSlash() {

	Echo.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))
}
