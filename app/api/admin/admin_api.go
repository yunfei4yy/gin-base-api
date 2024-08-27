package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"yunfei/app/request"
	"yunfei/app/response"
	"yunfei/app/service"
)

type adminApi struct{}

func (a *adminApi) Profile(c *gin.Context) {
	admin := service.AdminService.Profile()

	response.Success(c, admin)
}

func (a *adminApi) Save(c *gin.Context) {
	var admin request.SaveAdmin
	if err := c.ShouldBind(&admin); err != nil {
		fmt.Println(err)
		response.ValidateFailed(c, request.GetErrorMsg(admin, err))
		//response.ValidateFailed(c, err.Error())
	}
}
