package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// User struct
type User struct {
	ID              bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID          string        `json:"user_id" bson:"user_id"`
	FirstName       string        `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName        string        `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Email           string        `json:"email" bson:"email" validate:"required,email"`
	Password        string        `json:"password" bson:"password" validate:"required,min=6,max=100"`
	Role            string        `json:"role" bson:"role" validate:"oneof=ADMIN USER"`
	CreateAt        time.Time     `json:"create_at" bson:"create_at"`
	UpdateAt        time.Time     `json:"update_at" bson:"update_at"`
	Token           string        `json:"token" bson:"token"`
	RefreshToken    string        `json:"refresh_token" bson:"refresh_token"`
	FavouriteGenres []Genre       `json:"favourite_genres" bson:"favourite_genres" validate:"required,dive"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

// DTO
type UserResponse struct {
	UserID          string  `json:"user_id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	Token           string  `json:"token"`
	RefreshToken    string  `json:"refresh_token"`
	FavouriteGenres []Genre `json:"favourite_genres"`
}
