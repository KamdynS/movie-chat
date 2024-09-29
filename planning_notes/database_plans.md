# Database Plans for Room Management

## 1. Database Schema

### Rooms Table
```sql
CREATE TABLE rooms (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Room Members Table (for tracking users in rooms)
```sql
CREATE TABLE room_members (
    id SERIAL PRIMARY KEY,
    room_id VARCHAR(50) REFERENCES rooms(id),
    user_id BIGINT REFERENCES users(id),
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(room_id, user_id)
);
```

## 2. Go Structs

Update the `Room` struct in `internal/model/room.go`:

```go
type Room struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RoomMember struct {
    ID       int64     `json:"id" db:"id"`
    RoomID   string    `json:"room_id" db:"room_id"`
    UserID   int64     `json:"user_id" db:"user_id"`
    JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}
```

## 3. Repository Layer Updates

Update `internal/repository/room.go` to include the following methods:

```go
type RoomRepository interface {
    CreateRoom(ctx context.Context, room *model.Room) error
    GetRoom(ctx context.Context, id string) (*model.Room, error)
    GetRooms(ctx context.Context) ([]*model.Room, error)
    UpdateRoom(ctx context.Context, room *model.Room) error
    DeleteRoom(ctx context.Context, id string) error
    AddMember(ctx context.Context, roomID string, userID int64) error
    RemoveMember(ctx context.Context, roomID string, userID int64) error
    GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error)
}
```

Implement these methods using SQL queries.

## 4. Service Layer Updates

Update `internal/service/room.go` to use the new repository methods:

```go
type RoomService interface {
    CreateRoom(ctx context.Context, req *model.CreateRoomReq) (*model.Room, error)
    GetRoom(ctx context.Context, id string) (*model.Room, error)
    GetRooms(ctx context.Context) ([]*model.Room, error)
    UpdateRoom(ctx context.Context, room *model.Room) error
    DeleteRoom(ctx context.Context, id string) error
    JoinRoom(ctx context.Context, roomID string, userID int64) error
    LeaveRoom(ctx context.Context, roomID string, userID int64) error
    GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error)
}
```

## 5. Handler Layer Updates

Update `internal/handler/room.go` and `internal/handler/websocket.go` to use the new service methods.

## 6. Helper Function for Room ID Generation

Add a new file `internal/util/room_id.go`:

```go
package util

import (
    "crypto/rand"
    "encoding/base64"
)

func GenerateRoomID() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}
```

Use this function when creating new rooms.

## 7. Migration Script

Create a new migration file for adding the rooms and room_members tables:

```sql
-- Up
CREATE TABLE rooms (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE room_members (
    id SERIAL PRIMARY KEY,
    room_id VARCHAR(50) REFERENCES rooms(id),
    user_id BIGINT REFERENCES users(id),
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(room_id, user_id)
);

-- Down
DROP TABLE IF EXISTS room_members;
DROP TABLE IF EXISTS rooms;
```

## 8. WebSocket Hub Updates

Update the `Hub` struct in `internal/websocket/hub.go` to use room IDs from the database instead of storing Room objects directly:

```go
type Hub struct {
    Rooms      map[string]map[string]*Client // map[roomID]map[clientID]*Client
    Register   chan *Client
    Unregister chan *Client
    Broadcast  chan *Message
}
```

## 9. Testing

- Write unit tests for the new repository and service methods.
- Update existing tests to work with the new database-backed room system.
- Create integration tests to ensure the entire flow works correctly.

## 10. Data Migration (if needed)

If you have existing room data that needs to be migrated:
1. Create a script to read the current in-memory room data.
2. Insert this data into the new database tables.
3. Run this script as part of your deployment process.

## Implementation Steps

1. Create the new database tables.
2. Update the Go structs and interfaces.
3. Implement the new repository methods.
4. Update the service layer to use the new repository methods.
5. Update the handlers to use the new service methods.
6. Implement the room ID generation function.
7. Update the WebSocket hub to work with the new room system.
8. Write and run database migrations.
9. Update and create new tests.
10. Test thoroughly to ensure all functionality works as expected.
