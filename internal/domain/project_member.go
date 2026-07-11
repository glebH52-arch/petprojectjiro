package domain

import "time"

type Project_Members struct {
	Project_id int
	User_id    int
	Role       MemberRole
	JointedAt  time.Time
}

type MemberRole string

const (
	MemberRoleCreator MemberRole = "creator"
	MemberRoleAdmin   MemberRole = "admin"
	MemberRoleMember  MemberRole = "member"
)
