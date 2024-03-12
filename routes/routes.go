package routes

import (
	"ngetweet/controllers"

	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.Engine) {
	// Router users
	r.GET("/users", controllers.UserIndex)
	r.POST("/users", controllers.UserCreate)

	// Router tweets
	r.GET("/tweets", controllers.TweetIndex)
	r.POST("/tweets", controllers.TweetCreate)
}
