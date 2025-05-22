package model

import "encoding/json"

type Warehouse struct {
	ID        int    `db:"id" json:"id"`
	ShopID    int    `db:"shop_id" json:"shop_id"`
	Name      string `db:"name" json:"name"`
	IsActive  bool   `db:"is_active" json:"is_active"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type WarehouseWithStock struct {
	ID        int              `db:"id" json:"id"`
	ShopID    int              `db:"shop_id" json:"shop_id"`
	Name      string           `db:"name" json:"name"`
	IsActive  bool             `db:"is_active" json:"is_active"`
	CreatedAt string           `db:"created_at" json:"created_at"`
	UpdatedAt string           `db:"updated_at" json:"updated_at"`
	StockRaws json.RawMessage  `db:"stocks" json:"-"` // PENTING: db:"stocks"
	Stocks    []WarehouseStock `json:"stocks"`
}
