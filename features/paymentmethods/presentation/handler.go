package presentation

import (
	"bayareen-backend/features/paymentmethods"
	_payment_method_request "bayareen-backend/features/paymentmethods/presentation/request"
	_payment_method_response "bayareen-backend/features/paymentmethods/presentation/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	PaymentMethodBusiness paymentmethods.Business
}

func NewPaymentMethodHandler(paymentMethodBusiness paymentmethods.Business) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		PaymentMethodBusiness: paymentMethodBusiness,
	}
}

func (pmh *PaymentMethodHandler) Create(c echo.Context) error {
	pmRequest := _payment_method_request.PaymentMethod{}

	if err := c.Bind(&pmRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := pmh.PaymentMethodBusiness.Create(pmRequest.ToCore())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message:": "success",
		"data":     _payment_method_response.FromCore(resp),
	})

}

func (pmh *PaymentMethodHandler) GetAll(c echo.Context) error {
	resp := pmh.PaymentMethodBusiness.GetAll()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _payment_method_response.FromCoreSlice(resp),
	})
}

func (pmh *PaymentMethodHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := pmh.PaymentMethodBusiness.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _payment_method_response.FromCore(resp),
	})
}
