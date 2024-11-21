package repositories

import (
	// "encoding/json"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
	"tech-challenge-fase-1/internal/infra/database"
)

type OrderRepositoryDB struct {
	conn database.ConnectionDB
}

func NewOrderRepositoryDB(conn database.ConnectionDB) *OrderRepositoryDB {
	return &OrderRepositoryDB{conn: conn}
}

func (r *OrderRepositoryDB) Insert(order *entities.Order) error {
	sql := `
	INSERT INTO orders(id, customer_id, items, payment_status, preparation_status)
	VALUES ($1, $2, $3, $4, $5)
	`
	return r.conn.Exec(
		sql,
		order.GetId(),
		order.GetCustomerId(),
		newOrderItemHelperList(order.GetItems()),
		order.GetPaymentStatus().String(),
		order.GetPreparationStatus().String(),
	)
}

func (r *OrderRepositoryDB) FindOrderByID(orderId string) (*entities.Order, error) {
	sql := `
	SELECT
		id,
		customer_id,
		items,
		payment_status,
		preparation_status
	FROM orders 
	WHERE id = $1`
	row := r.conn.QueryRow(sql, orderId)
	return r.toEntity(row)
}

func (r *OrderRepositoryDB) Update(order *entities.Order) error {
	sql := `
	UPDATE orders 
	SET 
		customer_id = $1,
		items = $2, 
		payment_status = $3,
		preparation_status = $4
	WHERE id = $5;`
	return r.conn.Exec(
		sql,
		order.GetCustomerId(),
		newOrderItemHelperList(order.GetItems()),
		order.GetPaymentStatus().String(),
		order.GetPreparationStatus().String(),
		order.GetId(),
	)
}

func (r *OrderRepositoryDB) toEntity(row database.RowDB) (*entities.Order, error) {
	var id string
	var customerId *string
	var items []*orderItemHelper
	var paymentStatus entities.OrderPaymentStatus
	var preparationStatus entities.OrderPreparationStatus
	err := row.Scan(&id, &customerId, &items, &paymentStatus, &preparationStatus)
	if err != nil {
		if err.Error() == ErrNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return entities.RestoreOrder(
		id,
		customerId,
		orderItemsFromHelper(items),
		paymentStatus,
		preparationStatus,
	), nil
}

type orderItemHelper struct {
	Amount      float64  `json:"amount,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}

func orderItemsFromHelper(orderItemsHelper []*orderItemHelper) []*valueobjects.OrderItem {
	orderItems := make([]*valueobjects.OrderItem, 0)
	for _, item := range orderItemsHelper {
		orderItems = append(
			orderItems,
			valueobjects.NewOrderItem(item.Amount, item.Quantity, item.ProductName),
		)
	}
	return orderItems
}

func newOrderItemHelperList(orderItems []*valueobjects.OrderItem) []*orderItemHelper {
	orderItemsHelper := make([]*orderItemHelper, 0)
	for _, item := range orderItems {
		orderItemsHelper = append(
			orderItemsHelper,
			&orderItemHelper{
				Amount: item.GetAmount(),
				Quantity: item.GetQuantity(),
				ProductName: item.GetProductName(),
			},
		)
	}
	return orderItemsHelper
}
