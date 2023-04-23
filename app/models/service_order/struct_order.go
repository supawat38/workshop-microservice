package models

import (
	"time"

	"gorm.io/datatypes"
)

// รายการออเดอร์ (รายละเอียดหลัก)
type Orders struct {
	OrderCode     uint           `gorm:"primaryKey;autoIncrement;unique" json:"order_code" `
	MemberCode    uint           `json:"member_code"`
	Status        string         `json:"status" `
	Total         float64        `json:"total" `
	ProductDetail datatypes.JSON `json:"product_detail" gorm:"type:jsonb"`
	Remark        string         `json:"remark" `
	Created_at    time.Time      `json:"created_at"`
}

// สร้างใบสั้งซื้อ
type ReqOrder struct {
	MemberCode    uint           `json:"member_code"`
	ProductDetail datatypes.JSON `json:"product_detail" gorm:"type:jsonb"`
	Remark        string         `json:"remark" `
}

// รับค่า
type ReqOrderCode struct {
	OrderCode  uint `json:"order_code"`
	MemberCode uint `json:"member_code"`
}

// รายการออเดอร์ + สินค้า
type DetailOrder struct {
	OrderCode     uint           `json:"order_code" `
	MemberCode    uint           `json:"member_code"`
	MemberName    string         `json:"member_name"`
	Status        string         `json:"status" `
	Total         float64        `json:"total" `
	ProductDetail datatypes.JSON `json:"product_detail" gorm:"type:jsonb"`
	Remark        string         `json:"remark" `
	Created_at    time.Time      `json:"created_at"`
}
