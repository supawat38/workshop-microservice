package models

import "time"

// สำหรับ user
type Members struct {
	MemberCode uint      `gorm:"primaryKey;autoIncrement;unique" json:"member_code" validate:"required"`
	Username   string    `json:"username" `
	Password   string    `json:"password" `
	Name       string    `json:"name" `
	Status     int       `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
