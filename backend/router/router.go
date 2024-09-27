package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kamdyns/movie-chat/internal/user"
)

func InitRouter(userHandler *user.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	return r
}

func Start(addr string, r *gin.Engine) error {
	return r.Run(addr)
}
