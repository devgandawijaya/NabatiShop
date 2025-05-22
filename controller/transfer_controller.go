package controller

import (
	"fmt"
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransferController struct {
	Service *service.TransferService
}

func NewTransferController(service *service.TransferService) *TransferController {
	return &TransferController{Service: service}
}

func (c *TransferController) CreateTransfer(ctx echo.Context) error {
	var transfer model.WarehouseTransfer

	// Bind JSON ke struct
	if err := ctx.Bind(&transfer); err != nil {
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid input")
	}

	// Debug log
	fmt.Printf("Received Transfer from request: %+v\n", transfer)

	// Call service
	createdTransfer, err := c.Service.CreateTransfer(transfer)
	if err != nil {
		return utils.Respond(ctx, http.StatusInternalServerError, nil, err.Error())
	}

	return utils.Respond(ctx, http.StatusCreated, createdTransfer, "Transfer created successfully")
}

func (c *TransferController) GetAllTransfer(ctx echo.Context) error {
	shops, err := c.Service.GetAllTransfer()
	if err != nil {
		log.Printf("Error GetAllShop: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve shops")
	}
	return utils.Respond(ctx, http.StatusOK, shops, "Success")
}
