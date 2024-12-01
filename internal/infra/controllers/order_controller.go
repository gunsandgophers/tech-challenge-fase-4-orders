package controllers

import (
	"net/http"
	"tech-challenge-fase-1/internal/core/queries"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
	"tech-challenge-fase-1/internal/core/use_cases/orders"
	httpserver "tech-challenge-fase-1/internal/infra/http"
)

type OrderController struct {
	orderRepository       repositories.OrderRepositoryInterface
	customerService       services.CustomerServiceInterface
	productService        services.ProductServiceInterface
	paymentService        services.PaymentServiceInterface
	orderDisplayListQuery queries.OrderDisplayListQueryInterface
}

func NewOrderController(
	orderRepository repositories.OrderRepositoryInterface,
	customerService services.CustomerServiceInterface,
	productService services.ProductServiceInterface,
	paymentService services.PaymentServiceInterface,
	orderDisplayListQuery queries.OrderDisplayListQueryInterface,
) *OrderController {
	return &OrderController{
		orderRepository:       orderRepository,
		customerService:       customerService,
		productService:        productService,
		paymentService:        paymentService,
		orderDisplayListQuery: orderDisplayListQuery,
	}
}

// Checkout godoc
//
//	@Summary		Make an order checkout
//	@Description	make a checkout for an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			checkout	body		CheckoutRequest	true	"Checkout"
//	@Success		200			{object}	dtos.CheckoutDTO
//	@Failure		400			{string}	string	"when bad request"
//	@Failure		406			{string}	string	"when invalid params or invalid object"
//	@Router			/order/checkout [post]
func (cc *OrderController) Checkout(c httpserver.HTTPContext) {
	request := CheckoutRequest{}
	c.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	checkoutUseCase := orders.NewCheckoutOrderUseCase(
		cc.orderRepository,
		cc.customerService,
		cc.productService,
		cc.paymentService,
	)
	checkout, err := checkoutUseCase.Execute(request.CustomerId, request.ProductsIds)
	if err != nil {
		sendError(c, http.StatusNotAcceptable, err.Error())
		return
	}
	sendSuccess(c, http.StatusCreated, "checkout-order", checkout)
}

// OrderDisplayList godoc
//
//	@Summary		Get order list
//	@Description	Get order list for a display
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		dtos.OrderDisplayDTO
//	@Failure		400		{string}	string	"when bad request"
//	@Router			/order/display [get]
func (cc *OrderController) OrderDisplayList(c httpserver.HTTPContext) {
	orderDisplayListUseCase := orders.NewOrderDisplayListUseCase(
		cc.orderDisplayListQuery,
	)
	dtos, err := orderDisplayListUseCase.Execute()
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(c, http.StatusOK, "order-display-list", httpserver.Payload{
		"orders": dtos,
	})
}

// OrderPreparationStatusUpdate godoc
//
//	@Summary		Update order preparation status
//	@Description	Update the preparation status for an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id					path		string							true	"Order Identification"
//	@Param			preparation_status_update	body		PreparationStatusUpdateRequest	true	"Order Request Params"
//	@Success		204
//	@Failure		400							{string}	string	"when bad request"
//	@Router			/order/{order_id}/preparation-status [put]
func (cc *OrderController) OrderPreparationStatusUpdate(c httpserver.HTTPContext) {
	orderId := c.Param("order_id")
	request := &PreparationStatusUpdateRequest{}
	c.BindJSON(request)
	preparationStatusUpdateUseCase := orders.NewPreparationStatusUpdateUseCase(cc.orderRepository)
	err := preparationStatusUpdateUseCase.Execute(orderId, request.PreparationStatus)
	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	sendSuccess(c, http.StatusNoContent, "preparation-status-order", nil)
}
