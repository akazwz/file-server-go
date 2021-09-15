package initialize

import (
	"github.com/akazwz/file-server/model/response"
	"github.com/akazwz/file-server/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var router = gin.Default()

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	router.NoRoute(response.NotFound)
	router.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	publicRouter := router.Group("public")
	{
		routers.InitFileRouter(publicRouter)
	}

	return router
}
