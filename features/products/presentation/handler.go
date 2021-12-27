package presentation

import (
	"bayareen-backend/features/products"
	_product_request "bayareen-backend/features/products/presentation/request"
	_product_response "bayareen-backend/features/products/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductBusiness products.Business
}

func NewProductHandler(pb products.Business) *ProductHandler {
	return &ProductHandler{
		ProductBusiness: pb,
	}
}

func (ph *ProductHandler) Create(c echo.Context) error {
	productRequest := _product_request.Product{}
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ph.ProductBusiness.Create(productRequest.ToCore())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    _product_response.FromCore(resp),
	})
}