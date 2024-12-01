package orders

import (
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/core/entities"
	"tech-challenge-fase-1/internal/core/repositories"
	"tech-challenge-fase-1/internal/core/services"
)

type CheckoutOrderUseCase struct {
	orderRepository repositories.OrderRepositoryInterface
	customerService services.CustomerServiceInterface
	productService  services.ProductServiceInterface
	paymentService services.PaymentServiceInterface
}

func NewCheckoutOrderUseCase(
	orderRepository repositories.OrderRepositoryInterface,
	customerService services.CustomerServiceInterface,
	productService services.ProductServiceInterface,
	paymentService services.PaymentServiceInterface,
) *CheckoutOrderUseCase {
	return &CheckoutOrderUseCase{
		orderRepository: orderRepository,
		customerService: customerService,
		productService:  productService,
		paymentService:  paymentService,
	}
}

func (c *CheckoutOrderUseCase) validateCustomerId(customerId *string) error {
	if customerId == nil {
		return nil
	}
	_, err := c.customerService.GetCustomerById(*customerId)
	if err != nil {
		return err
	}
	return nil
}

func (c *CheckoutOrderUseCase) fetchProducts(productsIds []string) ([]*entities.Product, error) {
	products := []*entities.Product{}

	for _, productId := range productsIds {
		product, err := c.productService.FindProductByID(productId)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (c *CheckoutOrderUseCase) Execute(
	customerId *string,
	productsIds []string,
) (*dtos.CheckoutDTO, error) {
	if err := c.validateCustomerId(customerId); err != nil {
		return nil, err
	}
	products, err := c.fetchProducts(productsIds)
	if err != nil {
		return nil, err
	}
	order := entities.CreateOpenOrder(customerId)
	for _, product := range products {
		order.AddItem(product, 1)
	}

	// order.AwaitingPayment()
	// checkout, err := c.paymentGateway.Execute(
	// 	dtos.NewOrderDTOFromEntity(order),
	// 	dtos.PIX,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	checkout, err := c.paymentService.CreatePayment(order.GetId(), order.GetTotal())
	if err != nil {
		return nil, err
	}

	c.orderRepository.Insert(order)
	return checkout, nil
}
