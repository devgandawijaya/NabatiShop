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

type ShopController struct {
	Service *service.ShopService
}

func NewShopController(service *service.ShopService) *ShopController {
	return &ShopController{Service: service}
}

func (c *ShopController) GetAllShop(ctx echo.Context) error {
	shops, err := c.Service.GetAllShop()
	if err != nil {
		log.Printf("Error GetAllShop: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to retrieve shops")
	}
	return utils.Respond(ctx, http.StatusOK, shops, "Success")
}

func (c *ShopController) CreateShop(ctx echo.Context) error {
	var shop model.Shop
	if err := ctx.Bind(&shop); err != nil {
		log.Printf("Error binding CreateShop: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}

	result, err := c.Service.CreateShop(shop)
	if err != nil {
		log.Printf("Error CreateShop: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to create shop")
	}
	return utils.Respond(ctx, http.StatusCreated, result, "Shop created successfully")
}

func (c *ShopController) UpdateShop(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid shop ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid shop ID")
	}

	var shop model.Shop
	if err := ctx.Bind(&shop); err != nil {
		log.Printf("Error binding UpdateShop: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid request")
	}
	shop.ID = id

	result, err := c.Service.UpdateShop(shop)
	if err != nil {
		log.Printf("Error UpdateShop: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to update shop")
	}
	return utils.Respond(ctx, http.StatusOK, result, "Shop updated successfully")
}

func (c *ShopController) DeleteShop(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Invalid shop ID: %v", err)
		return utils.Respond(ctx, http.StatusBadRequest, nil, "Invalid shop ID")
	}

	if err := c.Service.DeleteShop(id); err != nil {
		log.Printf("Error DeleteShop: %v", err)
		return utils.Respond(ctx, http.StatusInternalServerError, nil, "Failed to delete shop")
	}
	return utils.Respond(ctx, http.StatusOK, nil, "Shop deleted successfully")
}
