package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kamdyns/movie-chat/internal/user"
	"github.com/kamdyns/movie-chat/internal/websocket"
)

func InitRouter(userHandler *user.Handler, wsHandler *websocket.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("ws/createRoom", wsHandler.CreateRoom)
	r.GET("ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("ws/getRooms", wsHandler.GetRooms)
	r.GET("ws/getClients/:roomId", wsHandler.GetClients)

	return r
}

func Start(addr string, r *gin.Engine) error {
	return r.Run(addr)
}
