package model

type WarehouseTransfer struct {
	ID              int    `db:"id" json:"id"`
	ProductID       int    `db:"product_id" json:"product_id"`
	FromWarehouseID int    `db:"from_warehouse_id" json:"from_warehouse_id"`
	ToWarehouseID   int    `db:"to_warehouse_id" json:"to_warehouse_id"`
	Quantity        int    `db:"quantity" json:"quantity"`
	TransferredAt   string `db:"transferred_at" json:"transferred_at"`
}
