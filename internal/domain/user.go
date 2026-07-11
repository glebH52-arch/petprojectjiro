package domain

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	AvatarURL *string
	Status    UserStatus
	CreatedAt time.Time
	DeletedAt *time.Time
}

type UserStatus string

const (
	UserStatusActive  UserStatus = "active"
	UserStatusDeleted UserStatus = "deleted"
)
