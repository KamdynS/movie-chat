package server

import (
	"database/sql"

	"github.com/clerkinc/clerk-sdk-go/clerk"
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
	clerkClient clerk.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)

	userService := service.NewUserService(userRepo)
	roomService := service.NewRoomService(roomRepo)

	wsHub := websocket.NewHub()

	clerkClient, err := clerk.NewClient(cfg.ClerkPublicKey)
	if err != nil {
		return nil, err
	}

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
		clerkClient: clerkClient,
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	userHandler := handler.NewUserHandler(s.userService)
	roomHandler := handler.NewRoomHandler(s.roomService)
	wsHandler := handler.NewWebSocketHandler(s.wsHub, s.roomService, s.userRepo)

	s.router.POST("/webhook", userHandler.HandleClerkWebhook)

	// Apply Clerk middleware to protected routes
	protected := s.router.Group("/")
	protected.Use(s.clerkAuthMiddleware())
	{
		protected.GET("/rooms", roomHandler.GetRooms)
		protected.GET("/ws", wsHandler.HandleWebSocket)
		protected.POST("/ws/createRoom", roomHandler.CreateRoom)
		protected.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
		protected.GET("/ws/getRooms", wsHandler.GetRooms)
		protected.GET("/ws/getClients/:roomId", wsHandler.GetClients)
	}
}

func (s *Server) clerkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken := c.GetHeader("Authorization")
		if sessionToken == "" {
			c.JSON(401, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		claims, err := s.clerkClient.VerifyToken(sessionToken)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid session token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.Subject)
		c.Next()
	}
}

func (s *Server) Run() error {
	go s.wsHub.Run()
	return s.router.Run(s.config.ServerAddress)
}
