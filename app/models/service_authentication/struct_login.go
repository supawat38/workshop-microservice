package models

// สำหรับ login
type Req_login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
