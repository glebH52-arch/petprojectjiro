package domain

import "time"

type Invite struct {
	ID           int
	ProjectID    int
	InviterID    int
	MemberID     int
	InviteStatus InviteStatus
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}

type InviteStatus string

const (
	InviteStatusPeding   InviteStatus = "peding"
	InviteStatusAccepted InviteStatus = "accepted"
	InviteStatusDecline  InviteStatus = "decline"
	InviteStatusExpired  InviteStatus = "expired"
)
