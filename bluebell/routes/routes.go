package routes

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"golearn/bluebell/controller"
	_ "golearn/bluebell/docs"
	"golearn/bluebell/logger"
	"golearn/bluebell/middlewares"
	"golearn/bluebell/setting"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/posts", controller.GetPostListHandler)
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.POST("/post/:id/reply", controller.CreateReplyHandler)
		v1.POST("/vote", controller.PostVoteController)

	}
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})
	r.NoRoute(func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"msg": "404"}) })
	return r
}
