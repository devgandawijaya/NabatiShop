package service

import (
	"go-api/model"
	"go-api/repository"
)

type WarehouseService struct {
	Repo *repository.WarehouseRepository
}

func NewWarehouseService(repo *repository.WarehouseRepository) *WarehouseService {
	return &WarehouseService{Repo: repo}
}

func (s *WarehouseService) GetAllWarehouse() ([]model.WarehouseWithStock, error) {
	return s.Repo.GetAllWarehouse()
}

// CreateWarehouse creates a new warehouse.
func (s *WarehouseService) CreateWarehouse(warehouse model.Warehouse) (*model.Warehouse, error) {
	return s.Repo.CreateWarehouses(warehouse)
}

func (s *WarehouseService) UpdateWarehouse(warehouse model.Warehouse) (*model.Warehouse, error) {
	return s.Repo.UpdateWarehouses(warehouse)
}

func (s *WarehouseService) DeleteWarehouse(id int) error {
	return s.Repo.DeleteWarehouses(id)
}
