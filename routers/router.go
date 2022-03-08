package routers

import (
	v1 "blog/api/v1"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.ServerMode)

	r := gin.Default()
	routerv1 := r.Group("api/v1")
	{
		// user
		routerv1.POST("user/add", v1.AddUser)
		routerv1.GET("users", v1.GetUserList)
		routerv1.PUT("user/:id", v1.UpdateUser)
		routerv1.DELETE("user/:id", v1.RemoveUser)
		// category
		// article
	}

	r.Run(utils.ServerPort)
}
