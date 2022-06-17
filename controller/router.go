package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine )  {
	api := r.Group("/api")
	{
		api.POST("/server/upload", Upload)
	}
}
