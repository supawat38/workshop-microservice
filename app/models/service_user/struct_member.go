package models

import "time"

// สำหรับ user Detail
type MembersDetail struct {
	MemberCode uint      `json:"member_code" validate:"required"`
	Username   string    `json:"username" `
	Password   string    `json:"password" `
	Name       string    `json:"name" `
	Status     int       `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
