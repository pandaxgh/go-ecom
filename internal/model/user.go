package model

import "time"

type User struct {
	Id         string    `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	UserName   string    `json:"userName"`
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"isVerfified"`
	IsDeleted  bool      `json:"isDeleted"`
	DeletedAt  time.Time `json:"deletedAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
