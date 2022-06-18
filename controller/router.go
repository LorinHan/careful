package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter(r *gin.Engine) {
	appMode := viper.GetString("App.Mode")
	api := r.Group("/api")
	{
		// 服务管理
		server := api.Group("/server")
		{
			server.POST("/upload", Server.Upload)

			if appMode == "main" {
				server.POST("", Server.Create)
				server.PUT("", Server.Update)
				server.DELETE("", Server.Delete)

				// 服务节点管理
				point := server.Group("/point")
				{
					point.GET("", ServerPoint.List)
					point.POST("", ServerPoint.Create)
					point.PUT("", ServerPoint.Update)
					point.DELETE("", ServerPoint.Delete)
				}
			}
		}

		if appMode == "main" {
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

		docker := api.Group("/docker")
		{
			docker.GET("list", Docker.List)
			docker.PUT("stop", Docker.Stop)
			docker.PUT("start", Docker.Start)
			docker.DELETE("", Docker.Delete)
			docker.GET("logs", Docker.Logs)
			docker.GET("info", Docker.Info)
			docker.POST("run", Docker.Run)
		}
	}
}
