package routes

import (
	"teste/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	v1 := router.Group("/v1")
	tweetController := controllers.NewTweetController()
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweets", tweetController.Create)
		v1.DELETE("/tweet/:id", tweetController.Delete)
	}
	return v1
}