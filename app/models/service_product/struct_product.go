package models

// สินค้า
type Products struct {
	ProductCode uint    `gorm:"primaryKey;autoIncrement;unique" json:"product_code" `
	ProductName string  `json:"product_name" `
	UnitPrice   float64 `json:"unit_price" `
	Stock       uint    `json:"stock" `
}
