package repository

import (
	"encoding/json"
	"fmt"
	"go-api/model"
	"go-api/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type WarehouseRepository struct {
	DB *sqlx.DB
}

func NewWarehouseRepository(db *sqlx.DB) *WarehouseRepository {
	return &WarehouseRepository{DB: db}
}

func (r *WarehouseRepository) GetAllWarehouse() ([]model.WarehouseWithStock, error) {
	var rawData []model.WarehouseWithStock

	err := r.DB.Select(&rawData, utils.GetAllWarehouseWithStock)
	if err != nil {
		log.Println("[ERROR] Failed to execute query:", err)
		return nil, err
	}

	for i, row := range rawData {
		fmt.Printf("Row #%d - Raw JSON: %s\n", i, string(row.StockRaws))

		var stocks []model.WarehouseStock
		if err := json.Unmarshal(row.StockRaws, &stocks); err != nil {
			log.Printf("[ERROR] Failed to unmarshal JSON for warehouse ID %d: %v\n", row.ID, err)
			return nil, err
		}
		rawData[i].Stocks = stocks
	}

	return rawData, nil
}

func (r *WarehouseRepository) CreateWarehouses(warehouse model.Warehouse) (*model.Warehouse, error) {
	var newWarehouse model.Warehouse

	err := r.DB.Get(&newWarehouse, utils.InsertWarehouse,
		warehouse.ShopID,
		warehouse.Name,
		warehouse.IsActive,
	)
	if err != nil {
		log.Printf("Error inserting warehouse: %v", err)
		return nil, err
	}

	return &newWarehouse, nil
}

func (r *WarehouseRepository) UpdateWarehouses(warehouse model.Warehouse) (*model.Warehouse, error) {
	var updatedWarehouse model.Warehouse

	err := r.DB.Get(&updatedWarehouse, utils.UpdateWarehouse,
		warehouse.Name,
		warehouse.IsActive,
		warehouse.ID,
	)
	if err != nil {
		log.Printf("Error updating warehouse: %v", err)
		return nil, err
	}

	return &updatedWarehouse, nil
}

func (r *WarehouseRepository) DeleteWarehouses(id int) error {
	_, err := r.DB.Exec(utils.DeleteWarehouse, id)
	if err != nil {
		log.Printf("Error deleting warehouse with id %d: %v", id, err)
	}
	return err
}
