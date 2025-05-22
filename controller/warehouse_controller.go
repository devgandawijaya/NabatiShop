package controller

import (
	"go-api/model"
	"go-api/service"
	"go-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WarehouseController struct {
	Service *service.WarehouseService
}

func NewWarehouseController(service *service.WarehouseService) *WarehouseController {
	return &WarehouseController{Service: service}
}

func (c *WarehouseController) GetAllWarehouses(ctx echo.Context) error {
	warehouses, err := c.Service.GetAllWarehouse()
	if err != nil {
		log.Printf("Error GetAllWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve warehouses")
	}
	return utils.Respond(ctx, http.StatusOK, warehouses, "Success")
}

func (c *WarehouseController) CreateWarehouses(ctx echo.Context) error {
	var warehouse model.Warehouse

	if err := ctx.Bind(&warehouse); err != nil {
		log.Printf("Error binding CreateWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}

	result, err := c.Service.CreateWarehouse(warehouse)
	if err != nil {
		log.Printf("Error CreateWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create warehouse")
	}

	return utils.Respond(ctx, http.StatusCreated, result, "Warehouse created successfully")
}

func (c *WarehouseController) UpdateWarehouses(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid warehouse ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid warehouse ID")
	}

	var warehouse model.Warehouse
	if err := ctx.Bind(&warehouse); err != nil {
		log.Printf("Error binding UpdateWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}
	warehouse.ID = id

	result, err := c.Service.UpdateWarehouse(warehouse)
	if err != nil {
		log.Printf("Error UpdateWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to update warehouse")
	}

	return utils.Respond(ctx, http.StatusOK, result, "Warehouse updated successfully")
}

func (c *WarehouseController) DeleteWarehouses(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid warehouse ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid warehouse ID")
	}

	if err := c.Service.DeleteWarehouse(id); err != nil {
		log.Printf("Error DeleteWarehouses: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete warehouse")
	}

	return utils.Respond(ctx, http.StatusOK, nil, "Warehouse deleted successfully")
}
