package service

import (
	"go-api/model"
	"go-api/repository"
)

type ShopService struct {
	Repo *repository.ShopRepository
}

func NewShopService(repo *repository.ShopRepository) *ShopService {
	return &ShopService{Repo: repo}
}

func (s *ShopService) GetAllShop() ([]model.ShopWithWarehouses, error) {
	return s.Repo.GetAllShop()
}

func (s *ShopService) CreateShop(shop model.Shop) (*model.Shop, error) {
	return s.Repo.CreateShop(shop)
}

func (s *ShopService) UpdateShop(shop model.Shop) (*model.Shop, error) {
	return s.Repo.UpdateShop(shop)
}

func (s *ShopService) DeleteShop(id int) error {
	return s.Repo.DeleteShop(id)
}
