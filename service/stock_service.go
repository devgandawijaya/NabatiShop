package service

import (
	"go-api/model"
	"go-api/repository"
)

type StockService struct {
	Repo *repository.StockRepository
}

func NewStockService(repo *repository.StockRepository) *StockService {
	return &StockService{Repo: repo}
}

func (s *StockService) GetStockByWarehouseID(id int) ([]model.WarehouseStock, error) {
	return s.Repo.GetStockByWarehouseID(id)
}

func (s *StockService) GetStockByWarehouseProductID(id int, product_id int) ([]model.WarehouseStock, error) {
	return s.Repo.GetStockByWarehouseProductID(id, product_id)
}

func (s *StockService) CreateStock(stock model.WarehouseStock) (*model.WarehouseStock, error) {
	return s.Repo.CreateStock(stock)
}

func (s *StockService) UpdateStock(stock model.WarehouseStock) (*model.WarehouseStock, error) {
	return s.Repo.UpdateStock(stock)
}
