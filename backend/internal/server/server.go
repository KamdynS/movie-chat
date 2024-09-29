package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/kamdyns/movie-chat/internal/config"
	"github.com/kamdyns/movie-chat/internal/handler"
	"github.com/kamdyns/movie-chat/internal/repository"
	"github.com/kamdyns/movie-chat/internal/service"
	"github.com/kamdyns/movie-chat/internal/websocket"
	"github.com/kamdyns/movie-chat/pkg/database"
)

type Server struct {
	config      *config.Config
	db          *sql.DB
	router      *gin.Engine
	userRepo    repository.UserRepository
	roomRepo    repository.RoomRepository
	userService service.UserService
	roomService service.RoomService
	wsHub       *websocket.Hub
}

func NewServer(cfg *config.Config) (*Server, error) {
	db, err := database.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)

	userService := service.NewUserService(userRepo)
	roomService := service.NewRoomService(roomRepo)

	wsHub := websocket.NewHub()

	router := gin.Default()

	server := &Server{
		config:      cfg,
		db:          db,
		router:      router,
		userRepo:    userRepo,
		roomRepo:    roomRepo,
		userService: userService,
		roomService: roomService,
		wsHub:       wsHub,
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	userHandler := handler.NewUserHandler(s.userService)
	roomHandler := handler.NewRoomHandler(s.roomService)
	wsHandler := handler.NewWebSocketHandler(s.wsHub)

	s.router.POST("/signup", userHandler.CreateUser)
	s.router.POST("/login", userHandler.Login)
	s.router.GET("/logout", userHandler.Logout)

	s.router.POST("/rooms", roomHandler.CreateRoom)
	s.router.GET("/rooms", roomHandler.GetRooms)

	s.router.GET("/ws", wsHandler.HandleWebSocket)
	s.router.POST("/ws/createRoom", wsHandler.CreateRoom)
	s.router.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	s.router.GET("/ws/getRooms", wsHandler.GetRooms)
	s.router.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func (s *Server) Run() error {
	go s.wsHub.Run()
	return s.router.Run(s.config.ServerAddress)
}
