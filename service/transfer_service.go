package service

import (
	"go-api/model"
	"go-api/repository"
)

type TransferService struct {
	Repo *repository.TransferRepository
}

func NewTransferService(repo *repository.TransferRepository) *TransferService {
	return &TransferService{Repo: repo}
}

func (s *TransferService) CreateTransfer(transfer model.WarehouseTransfer) (*model.WarehouseTransfer, error) {
	return s.Repo.CreateTransfer(transfer)
}

func (s *TransferService) GetAllTransfer() ([]model.WarehouseTransfer, error) {
	return s.Repo.GetAllTransfer()
}
