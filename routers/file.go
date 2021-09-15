package routers

import (
	"github.com/akazwz/file-server/api/file"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(Router *gin.RouterGroup) {
	FileRouter := Router.Group("file")
	{
		FileRouter.POST("", file.CreateFile)
	}
}
