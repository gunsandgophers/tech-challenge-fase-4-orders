package queries

import (
	"log"
	"tech-challenge-fase-1/internal/core/dtos"
	"tech-challenge-fase-1/internal/infra/database"
	"time"
)

type OrderDisplayListQueryDB struct {
	conn database.ConnectionDB
}

func NewOrderDisplayListQueryDB(conn database.ConnectionDB) *OrderDisplayListQueryDB {
	return &OrderDisplayListQueryDB{conn: conn}
}

func (q *OrderDisplayListQueryDB) Execute() ([]*dtos.OrderDisplayDTO, error) {
	sql := `
	SELECT 
		id,
		customer_id,
		items,
		preparation_status,
		created_at,
		CASE preparation_status
			WHEN 'READY' THEN 0
			WHEN 'IN_PREPARATION' THEN 1
			ELSE 2 
		END preparation_order_by
	FROM public.orders
	WHERE
		preparation_status NOT IN ('AWAITING','CANCELED','FINISHED')
	ORDER BY
		preparation_order_by, created_at
	`
	rows, err := q.conn.Query(sql)
	if err != nil {
		return nil, err
	}
	return q.toDTOList(rows), nil
}

func (q *OrderDisplayListQueryDB) toDTOList(rows database.RowsDB) []*dtos.OrderDisplayDTO {
	var orders []*dtos.OrderDisplayDTO
	for rows.Next() {
		if o, err := q.toDTO(rows); err == nil {
			orders = append(orders, o)
		} else {
			log.Println(err)
		}
	}
	return orders
}

func (q *OrderDisplayListQueryDB) toDTO(
	row database.RowDB,
) (*dtos.OrderDisplayDTO, error) {
	var id string
	var customerId *string
	var items []*orderItemDisplayHelper
	var preparationStatus string
	var createdAt time.Time
	var preparationOrderBy int

	err := row.Scan(
		&id,
		&customerId,
		&items,
		&preparationStatus,
		&createdAt,
		&preparationOrderBy,
	)
	if err != nil {
		return nil, err
	}

	orderItemsDisplay := make([]*dtos.OrderItemDisplayDTO, 0)
	for _, item := range items {
		orderItemsDisplay = append(orderItemsDisplay, &dtos.OrderItemDisplayDTO{
			Quantity: item.Quantity,
			ProductName: item.ProductName,
		})
	}

	return &dtos.OrderDisplayDTO{
		Id: id,
		CustomerId: customerId,
		Items: orderItemsDisplay,
		PreparationStatus: preparationStatus,
		CreatedAt: createdAt,
	}, nil
}

type orderItemDisplayHelper struct {
	Amount      float64  `json:"amount,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	ProductName string `json:"product_name,omitempty"`
}
