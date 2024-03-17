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
	r.POST("/logout", controllers.Logout)

	// Router users
	r.GET("/users", middleware.RequiredAuth, controllers.UserIndex)

	// Router tweets
	r.GET("/tweets", middleware.RequiredAuth, controllers.TweetIndex)
	r.POST("/tweets", middleware.RequiredAuth, controllers.TweetCreate)
	r.DELETE("/tweets/:id/tweet", middleware.RequiredAuth, controllers.DeleteTweet)

	// Router Likes
	r.GET("/likes", middleware.RequiredAuth, controllers.LikeIndex)
	r.PUT("/tweets/:id/like", middleware.RequiredAuth, controllers.AddLike)

	// Router Comments
	r.GET("/comments", middleware.RequiredAuth, controllers.CommentIndex)
	r.POST("/tweets/:id/comment", middleware.RequiredAuth, controllers.AddComment)
	r.DELETE("/tweets/:id/comment", middleware.RequiredAuth, controllers.DeleteComment)
}
