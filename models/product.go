package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ProductID     uint   `json:"product_id"`
	ProductName   string `json:"product_name"`
	BarcodeA      string `json:"barcode_a"`
	BarcodeB      string `json:"barcode_b"`
	QrCode        string `json:"qrcode"`
	Weight        int    `json:"weight"`
	Amount        int    `json:"amount"`
	AmountChecked int    `json:"amount_checked"`

	TicketID uint `json:"ticket_id"`
}
