## 1. Currently implemented

### User Management
- POST /signup: Create a new user account
- POST /login: Authenticate a user and return a JWT
- GET /logout: Log out a user (invalidate JWT)

### Room Management
- POST /rooms: Create a new chat room
- GET /rooms: Get a list of all chat rooms
- GET /rooms/{id}: Get details of a specific chat room
- PUT /rooms/{id}: Update a chat room's details
- DELETE /rooms/{id}: Delete a chat room
- POST /rooms/{id}/members: Add a member to a chat room
- DELETE /rooms/{id}/members/{user_id}: Remove a member from a chat room
- GET /rooms/{id}/members: Get a list of members in a chat room

### WebSocket
- GET /ws: WebSocket endpoint for real-time chat functionality

## 2. To be implemented

### TV Guide
- GET /tv-guide: Fetch the TV Guide data
- GET /tv-guide/{id}: Fetch a specific TV Guide entry
- POST /tv-guide: Add a new TV Guide entry
- PUT /tv-guide/{id}: Update a TV Guide entry
- DELETE /tv-guide/{id}: Delete a TV Guide entry

### Enhanced Room Management
- POST /rooms/{id}/schedule: Schedule a room for a specific time
- GET /rooms/scheduled: Get all scheduled rooms
- POST /rooms/{id}/categories: Add categories to a room
- POST /rooms/{id}/tags: Add tags to a room
- GET /rooms/categories: Get rooms by category
- GET /rooms/tags: Get rooms by tag
- GET /rooms/trending: Get trending rooms based on user activity

### User Profile
- GET /users/{id}/profile: Get user profile information
- PUT /users/{id}/profile: Update user profile information
- GET /users/{id}/history: Get user's chat history

### Moderation
- POST /reports: Submit a report for inappropriate content
- GET /reports: Get all reports (for moderators)
- PUT /reports/{id}: Update report status

### Authentication (Clerk Integration)
- Remove existing auth endpoints (/signup, /login, /logout)
- Implement Clerk webhook endpoint for user creation/updates
- POST /clerk/webhook: Handle Clerk authentication events

### Movie Data Integration
- GET /movies: Get a list of movies (potentially with filtering)
- GET /movies/{id}: Get details of a specific movie
- GET /movies/search: Search for movies by title, genre, etc.

### Recommendations
- GET /users/{id}/recommendations: Get movie or room recommendations for a user

### Analytics
- GET /analytics/rooms: Get analytics data for chat rooms
- GET /analytics/users: Get user engagement analytics

### Admin
- GET /admin/users: Get a list of all users (for admins)
- PUT /admin/users/{id}: Update user roles or permissions
- GET /admin/rooms: Get a list of all rooms with additional details
