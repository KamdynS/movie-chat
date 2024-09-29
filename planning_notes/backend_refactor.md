# Backend Refactoring Guide

## 1. Directory Structure

Reorganize your backend code into a more modular structure:

```
.
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── auth
│   │   └── middleware.go
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   ├── room.go
│   │   ├── user.go
│   │   └── websocket.go
│   ├── model
│   │   ├── room.go
│   │   └── user.go
│   ├── repository
│   │   ├── room.go
│   │   └── user.go
│   ├── service
│   │   ├── room.go
│   │   └── user.go
│   └── websocket
│       ├── client.go
│       └── hub.go
├── pkg
│   └── database
│       └── postgres.go
└── go.mod
```

## 2. Implement Dependency Injection

Use a dependency injection pattern to manage dependencies:

1. Create a `Server` struct to hold all dependencies:

```go
// internal/server/server.go
type Server struct {
    config     *config.Config
    db         *sql.DB
    router     *gin.Engine
    userRepo   repository.UserRepository
    roomRepo   repository.RoomRepository
    userService service.UserService
    roomService service.RoomService
}

func NewServer(config *config.Config) (*Server, error) {
    // Initialize dependencies
    // Return new Server instance
}

func (s *Server) Run() error {
    return s.router.Run(s.config.ServerAddress)
}
```

2. Update `main.go` to use the new `Server` struct:

```go
// cmd/server/main.go
func main() {
    config, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    server, err := server.NewServer(config)
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }

    if err := server.Run(); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
```

## 3. Use Interfaces for Repositories and Services

Define interfaces for your repositories and services:

```go
// internal/repository/user.go
type UserRepository interface {
    CreateUser(ctx context.Context, user *model.User) error
    GetUserByID(ctx context.Context, id string) (*model.User, error)
    // Add other methods as needed
}

// internal/service/user.go
type UserService interface {
    CreateUser(ctx context.Context, user *model.User) error
    GetUserByID(ctx context.Context, id string) (*model.User, error)
    // Add other methods as needed
}
```

Implement these interfaces in separate files:

```go
// internal/repository/user_postgres.go
type userPostgresRepo struct {
    db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) UserRepository {
    return &userPostgresRepo{db: db}
}

// Implement methods...

// internal/service/user_service.go
type userService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
    return &userService{repo: repo}
}

// Implement methods...
```

## 4. Use Context for Request Scoping

Ensure all repository and service methods accept a `context.Context` as their first parameter:

```go
func (r *userPostgresRepo) GetUserByID(ctx context.Context, id string) (*model.User, error) {
    // Use ctx when querying the database
    row := r.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id)
    // ...
}
```

## 5. Implement Proper Error Handling

Create custom error types and use them consistently:

```go
// pkg/errors/errors.go
type ErrorType string

const (
    ErrorTypeNotFound ErrorType = "NOT_FOUND"
    ErrorTypeInvalidInput ErrorType = "INVALID_INPUT"
    // Add other error types as needed
)

type Error struct {
    Type    ErrorType
    Message string
}

func (e *Error) Error() string {
    return e.Message
}

func NewError(errorType ErrorType, message string) *Error {
    return &Error{
        Type:    errorType,
        Message: message,
    }
}
```

Use these custom errors in your code:

```go
if user == nil {
    return nil, NewError(ErrorTypeNotFound, "user not found")
}
```

## 6. Implement Logging

Use a structured logging library like `zap` for better log management:

```go
// pkg/logger/logger.go
import "go.uber.org/zap"

var log *zap.Logger

func Init() {
    var err error
    log, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
}

func Info(msg string, fields ...zap.Field) {
    log.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    log.Error(msg, fields...)
}

// Add other log levels as needed
```

Use this logger throughout your application:

```go
logger.Info("User created", zap.String("userId", user.ID))
```

## 7. Implement Graceful Shutdown

Add graceful shutdown to your server:

```go
// internal/server/server.go
func (s *Server) Run() error {
    srv := &http.Server{
        Addr:    s.config.ServerAddress,
        Handler: s.router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exiting")
    return nil
}
```

These refactoring steps will help improve the organization, maintainability, and scalability of your backend code. Remember to adjust and implement these changes gradually, testing each step to ensure everything works as expected.