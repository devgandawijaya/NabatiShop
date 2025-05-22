package repository

import (
	"fmt"
	"go-api/model"
	"go-api/utils"

	"github.com/jmoiron/sqlx"
)

type TransferRepository struct {
	DB *sqlx.DB
}

func NewTransferRepository(db *sqlx.DB) *TransferRepository {
	return &TransferRepository{DB: db}
}

func (r *TransferRepository) CreateTransfer(transfer model.WarehouseTransfer) (*model.WarehouseTransfer, error) {
	var isActiveFrom, isActiveTo bool

	err := r.DB.Get(&isActiveFrom, utils.CheckWarehouseActive, transfer.FromWarehouseID)
	if err != nil || !isActiveFrom {
		return nil, fmt.Errorf("gudang asal tidak aktif")
	}

	err = r.DB.Get(&isActiveTo, utils.CheckWarehouseActive, transfer.ToWarehouseID)
	if err != nil || !isActiveTo {
		return nil, fmt.Errorf("gudang tujuan tidak aktif")
	}

	var fromStock model.WarehouseStock
	err = r.DB.Get(&fromStock, utils.GetStock, transfer.FromWarehouseID, transfer.ProductID)
	if err != nil {
		return nil, fmt.Errorf("stok tidak ditemukan di gudang asal")
	}

	if fromStock.AvailableQty < transfer.Quantity {
		return nil, fmt.Errorf("stok gudang asal tidak mencukupi")
	}

	newQtyFrom := fromStock.AvailableQty - transfer.Quantity
	_, err = r.DB.Exec(utils.UpdateStock, newQtyFrom, fromStock.ID)
	if err != nil {
		return nil, fmt.Errorf("gagal mengurangi stok dari gudang asal")
	}

	var toStock model.WarehouseStock
	err = r.DB.Get(&toStock, utils.GetStock, transfer.ToWarehouseID, transfer.ProductID)
	if err != nil {
		insertQuery := `
			INSERT INTO warehouse_stocks (warehouse_id, product_id, available_qty, reserved_qty)
			VALUES ($1, $2, $3, 0)
		`
		_, err = r.DB.Exec(insertQuery, transfer.ToWarehouseID, transfer.ProductID, transfer.Quantity)
		if err != nil {
			return nil, fmt.Errorf("gagal menambahkan stok ke gudang tujuan")
		}
	} else {
		newQtyTo := toStock.AvailableQty + transfer.Quantity
		_, err = r.DB.Exec(utils.UpdateStock, newQtyTo, toStock.ID)
		if err != nil {
			return nil, fmt.Errorf("gagal menambahkan stok ke gudang tujuan")
		}
	}

	var createdTransfer model.WarehouseTransfer
	err = r.DB.Get(&createdTransfer, utils.InsertTransfer,
		transfer.ProductID,
		transfer.FromWarehouseID,
		transfer.ToWarehouseID,
		transfer.Quantity,
	)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan data transfer: %v", err)
	}

	return &createdTransfer, nil
}

func (r *TransferRepository) GetAllTransfer() ([]model.WarehouseTransfer, error) {
	var transfers []model.WarehouseTransfer

	err := r.DB.Select(&transfers, utils.GetAllTransfers)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}
