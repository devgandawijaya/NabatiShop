package repository

import (
	"fmt"
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type StockRepository struct {
	DB *sqlx.DB
}

func NewStockRepository(db *sqlx.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) GetStockByWarehouseID(id int) ([]model.WarehouseStock, error) {
	var stocks []model.WarehouseStock

	err := r.DB.Select(&stocks, utils.GetStockByWarehouseID, id)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepository) GetStockByWarehouseProductID(id int, product_id int) ([]model.WarehouseStock, error) {
	var stocks []model.WarehouseStock

	err := r.DB.Select(&stocks, utils.GetStockByWarehouseProductID, id, product_id)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepository) CreateStock(stock model.WarehouseStock) (*model.WarehouseStock, error) {
	query := `
		INSERT INTO warehouse_stocks (warehouse_id, product_id, available_qty, reserved_qty)
		VALUES ($1, $2, $3, $4)
		RETURNING id, warehouse_id, product_id, available_qty, reserved_qty
	`

	fmt.Printf("Executing query: %s\n", query)
	fmt.Printf("Params: warehouse_id=%d, product_id=%d, available_qty=%d, reserved_qty=%d\n",
		stock.WarehouseID, stock.ProductID, stock.AvailableQty, stock.ReservedQty)

	var createdStock model.WarehouseStock
	err := r.DB.Get(&createdStock, query, stock.WarehouseID, stock.ProductID, stock.AvailableQty, stock.ReservedQty)
	if err != nil {
		return nil, err
	}

	return &createdStock, nil
}

func (r *StockRepository) UpdateStock(stock model.WarehouseStock) (*model.WarehouseStock, error) {

	query := `
		UPDATE warehouse_stocks
		SET available_qty = $1
		WHERE id = $2 AND warehouse_id = $3 AND product_id = $4
		RETURNING id, warehouse_id, product_id, available_qty, reserved_qty
	`

	fmt.Printf("Executing query: %s\n", query)
	fmt.Printf("Params: available_qty=%d, id=%d, warehouse_id=%d, product_id=%d\n",
		stock.AvailableQty, stock.ID, stock.WarehouseID, stock.ProductID)

	var updatedStock model.WarehouseStock
	err := r.DB.Get(&updatedStock, query, stock.AvailableQty, stock.ID, stock.WarehouseID, stock.ProductID)
	if err != nil {
		return nil, err
	}
	return &updatedStock, nil
}
