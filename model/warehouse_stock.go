package model

type WarehouseStock struct {
	ID           int `db:"id" json:"id"`
	WarehouseID  int `db:"warehouse_id" json:"warehouse_id"`
	ProductID    int `db:"product_id" json:"product_id"`
	AvailableQty int `db:"available_qty" json:"available_qty"`
	ReservedQty  int `db:"reserved_qty" json:"reserved_qty"`
}
