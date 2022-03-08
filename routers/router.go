package routers

import (
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.ServerMode)

	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("/hello", func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "hello world",
				},
			)
		})
	}

	r.Run(utils.ServerPort)
}
