package model

type User struct {
	ID          int64  `json:"id" db:"id"`
	ClerkUserID string `json:"clerk_user_id" db:"clerk_user_id"`
	Username    string `json:"username" db:"username"`
	Email       string `json:"email" db:"email"`
}

// Remove CreateUserReq, CreateUserRes, LoginUserReq, LoginUserRes

type ClerkWebhookEvent struct {
	Type string `json:"type"`
	Data struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email_address"`
	} `json:"data"`
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
