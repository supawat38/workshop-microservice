package models

import (
	"time"

	"gorm.io/datatypes"
)

// รายการออเดอร์ + สินค้า
type DetailOrderByMember struct {
	OrderCode     uint           `json:"order_code" `
	MemberCode    uint           `json:"member_code"`
	MemberName    string         `json:"member_name"`
	Status        string         `json:"status" `
	Total         float64        `json:"total" `
	ProductDetail datatypes.JSON `json:"product_detail" gorm:"type:jsonb"`
	Remark        string         `json:"remark" `
	Created_at    time.Time      `json:"created_at"`
}
