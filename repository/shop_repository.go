package repository

import (
	"encoding/json"
	"go-api/model"
	"go-api/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type ShopRepository struct {
	DB *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *ShopRepository {
	return &ShopRepository{DB: db}
}

func (r *ShopRepository) GetAllShop() ([]model.ShopWithWarehouses, error) {

	var rawData []model.ShopWithWarehouses
	err := r.DB.Select(&rawData, utils.GetAllShopActive)
	if err != nil {
		return nil, err
	}

	for i, row := range rawData {
		var warehouses []model.Warehouse
		if err := json.Unmarshal(row.WarehousesRaw, &warehouses); err != nil {
			return nil, err
		}
		rawData[i].Warehouses = warehouses
	}
	return rawData, nil

}

func (r *ShopRepository) CreateShop(shop model.Shop) (*model.Shop, error) {
	var newShop model.Shop
	err := r.DB.Get(&newShop, utils.InsertShop, shop.Name)
	if err != nil {
		log.Printf("Error inserting shop: %v", err)
		return nil, err
	}
	return &newShop, nil
}

func (r *ShopRepository) UpdateShop(shop model.Shop) (*model.Shop, error) {
	query := `
		UPDATE shops
		SET name = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id, name, created_at, updated_at
	`
	var updatedShop model.Shop
	err := r.DB.Get(&updatedShop, query, shop.Name, shop.ID)
	return &updatedShop, err
}

func (r *ShopRepository) DeleteShop(id int) error {
	query := `DELETE FROM shops WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
