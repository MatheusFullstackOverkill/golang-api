package interfaces

import (
	"time"
)

type UserRequest struct {
	UserID      int    `json:"user_id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	UserPicture string `json:"user_picture"`
	Password    string `form:"password"`
}

type UserResponse struct {
	UserID      int       `json:"user_id"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	UserPicture string    `json:"user_picture"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	UserID      int       `json:"user_id"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	UserPicture string    `json:"user_picture"`
	Password    string    `form:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type ListParams struct {
	Offset string `query:"offset"`
	Limit  string `query:"limit"`
}

type ListResult struct {
	Data  []interface{} `json:"data"`
	Count int           `json:"count"`
}
