package route

import (
	"github.com/gin-gonic/gin"
	"yunfei/app/api/admin"
	"yunfei/app/middleware"
)

func genAdminRouter(rg *gin.RouterGroup) {
	rg.Use(middleware.CorsMiddleware)
	// 登录接口
	rg.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "login"})
	})
	// ----以上接口不走权限校验----

	// 注册权限验证中间件
	rg.Use(middleware.PermissionMiddleware)
	//rg.GET("/a", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"message": "a"})
	//})
	rg.GET("/profile", admin.AdminApi.Profile)
	rg.POST("/save", admin.AdminApi.Save)
}
