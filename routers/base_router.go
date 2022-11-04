
package routers

import (
	"github.com/gin-gonic/gin"
	"haotian_main/controller"
)

const (
	DemoROOT = "/main/api/user"
	SellerROOT = "/main/seller"
)

//路由设置
func RegisterRouter(router *gin.Engine) {
	routerUser(router)
	routerSeller(router)
}
//用户路由
func routerUser(engine *gin.Engine) {
	var group = engine.Group(DemoROOT)
	{
		con := &controller.UserController{}
		group.POST("/add", con.Add)
		group.GET("/get", con.Get)
		group.GET("/get_grpc", con.GetGrpc)
	}
}

func routerSeller(engine *gin.Engine) {
	var group = engine.Group(SellerROOT)
	{
		con := &controller.SellerController{}
		group.POST("/bind_member", con.BindMember)
	}
}