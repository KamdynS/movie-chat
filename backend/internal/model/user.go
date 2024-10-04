package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ClerkUserID string    `json:"clerk_user_id" db:"clerk_user_id"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

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
