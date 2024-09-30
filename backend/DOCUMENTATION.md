# API Documentation

## Authentication

All protected routes require a Clerk session token in the Authorization header:

Authorization: YOUR_CLERK_SESSION_TOKEN

## WebSocket Connection

- **URL:** `/ws`
- **Method:** `GET`
- **Query Parameters:**
  - `roomId`: string
- **Response:** WebSocket connection

## Room Endpoints

### Create Room

- **URL:** `/ws/createRoom`
- **Method:** `GET`
- **Query Parameters:**
  - `name`: string
- **Response:**
  ```json
  {
    "id": "string",
    "name": "string"
  }
  ```

### Get Rooms

- **URL:** `/ws/getRooms`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": "string",
      "name": "string"
    }
  ]
  ```

### Get Clients in Room

- **URL:** `/ws/getClients/:roomId`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": "string",
      "username": "string"
    }
  ]
  ```

### Get Rooms (HTTP)

- **URL:** `/rooms`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": "string",
      "name": "string"
    }
  ]
  ```

### Get Room

- **URL:** `/rooms/:id`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "id": "string",
    "name": "string"
  }
  ```

### Update Room

- **URL:** `/rooms/:id`
- **Method:** `PUT`
- **Body:**
  ```json
  {
    "name": "string"
  }
  ```
- **Response:**
  ```json
  {
    "id": "string",
    "name": "string"
  }
  ```

### Delete Room

- **URL:** `/rooms/:id`
- **Method:** `DELETE`
- **Response:** 200 OK

### Add Member to Room

- **URL:** `/rooms/:id/members`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "user_id": 123
  }
  ```
- **Response:** 200 OK

### Remove Member from Room

- **URL:** `/rooms/:id/members/:user_id`
- **Method:** `DELETE`
- **Response:** 200 OK

### Get Room Members

- **URL:** `/rooms/:id/members`
- **Method:** `GET`
- **Response:**
  ```json
  [
    {
      "id": 123,
      "room_id": "string",
      "user_id": 456,
      "joined_at": "2023-04-20T12:00:00Z"
    }
  ]
  ```

## User Endpoints

### Handle Clerk Webhook

- **URL:** `/webhook`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "type": "user.created",
    "data": {
      "id": "clerk_user_123",
      "username": "string",
      "email_address": "string"
    }
  }
  ```
- **Response:** 200 OK

## WebSocket Messages

### Incoming Messages

- **Send Message:**
  ```json
  {
    "content": "string",
    "room_id": "string",
    "username": "string"
  }
  ```

### Outgoing Messages

- **New User Joined:**
  ```json
  {
    "content": "A new user has joined the room",
    "room_id": "string",
    "username": "string"
  }
  ```

- **User Left:**
  ```json
  {
    "content": "User has left the room",
    "room_id": "string",
    "username": "string"
  }
  ```

- **New Message:**
  ```json
  {
    "content": "string",
    "room_id": "string",
    "username": "string"
  }
  ```