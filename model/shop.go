package model

import "encoding/json"

type Shop struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type ShopWithWarehouses struct {
	ID            int             `db:"id" json:"id"`
	Name          string          `db:"name" json:"name"`
	CreatedAt     string          `db:"created_at" json:"created_at"`
	UpdatedAt     string          `db:"updated_at" json:"updated_at"`
	WarehousesRaw json.RawMessage `db:"warehouses" json:"-"`
	Warehouses    []Warehouse     `json:"warehouses"`
}
