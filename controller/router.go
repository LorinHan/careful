package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		server := api.Group("/server")
		{
			server.POST("/upload", Upload)
		}

		folder := api.Group("/folder")
		{
			folder.POST("", Folder.Create)
			folder.GET("", Folder.List)
			folder.PUT("", Folder.Update)
			folder.DELETE("", Folder.Delete)
		}
	}
}
