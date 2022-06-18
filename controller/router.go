package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		// 服务管理
		server := api.Group("/server")
		{
			server.POST("/upload", Upload)
		}

		// 文件夹管理
		folder := api.Group("/folder")
		{
			folder.POST("", Folder.Create)
			folder.GET("", Folder.List)
			folder.PUT("", Folder.Update)
			folder.DELETE("", Folder.Delete)
		}

		// 节点管理
		node := api.Group("/node")
		{
			node.POST("", Node.Create)
			node.GET("", Node.List)
			node.PUT("", Node.Update)
			node.DELETE("", Node.Delete)
		}
	}
}
