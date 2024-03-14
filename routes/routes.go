package routes

import (
	"ngetweet/controllers"
	"ngetweet/middleware"

	"github.com/gin-gonic/gin"
)

func RouteInit(r *gin.Engine) {
	// Router Auth
	r.POST("/users", controllers.Register)
	r.POST("/login", controllers.Login)

	// Router users
	r.GET("/users", controllers.UserIndex)

	// Router tweets
	r.GET("/tweets", middleware.RequiredAuth, controllers.TweetIndex)
	r.POST("/tweets", middleware.RequiredAuth, controllers.TweetCreate)
	r.PUT("/tweets/:id/like", middleware.RequiredAuth, controllers.AddLike)

	// Router Likes
	r.GET("/likes", middleware.RequiredAuth, controllers.LikeIndex)

	// Routes Following
	r.POST("/follow",controllers.Following)
	r.GET("/follow",controllers.FollowIndex)

	// Routes Followers
	r.GET("/followers",controllers.FollowersIndex)
}
