package repositories

import (
	// "encoding/json"
	"context"
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoOrderItem struct {
	Amount      float64
	Quantity    int
	ProductName string
}

type mongoOrder struct {
	ID                string           `json:"_id,omitempty"`
	CustomerId        *string          `json:"customer_id,omitempty"`
	Items             []mongoOrderItem `json:"items,omitempty"`
	PaymentStatus     string           `json:"payment_status,omitempty"`
	PreparationStatus string           `json:"preparation_status,omitempty"`
}

func toMongoOrder(order *entities.Order) *mongoOrder {
	items := make([]mongoOrderItem, 0)
	for _, i := range order.GetItems() {
		items = append(items, mongoOrderItem{
			Amount:      i.GetAmount(),
			Quantity:    i.GetQuantity(),
			ProductName: i.GetProductName(),
		})
	}

	return &mongoOrder{
		ID:                order.GetId(),
		CustomerId:        order.GetCustomerId(),
		Items:             items,
		PaymentStatus:     order.GetPaymentStatus().String(),
		PreparationStatus: order.GetPreparationStatus().String(),
	}
}

func toOrder(mOrder *mongoOrder) *entities.Order {
	items := make([]*valueobjects.OrderItem, 0)
	for _, i := range mOrder.Items {
		items = append(items, valueobjects.NewOrderItem(i.Amount, i.Quantity, i.ProductName))
	}
	return entities.RestoreOrder(
		mOrder.ID,
		mOrder.CustomerId,
		items,
		entities.OrderPaymentStatus(mOrder.PaymentStatus),
		entities.OrderPreparationStatus(mOrder.PreparationStatus),
	)
}

type OrderRepositoryMongo struct {
	db *mongo.Database
}

const orderCollectionName = "orders"

func NewOrderRepositoryMongo(db *mongo.Database) *OrderRepositoryMongo {
	return &OrderRepositoryMongo{db: db}
}

func (r *OrderRepositoryMongo) Insert(order *entities.Order) error {
	collection := r.db.Collection(orderCollectionName)
	mongoOrder := toMongoOrder(order)
	_, err := collection.InsertOne(context.TODO(), mongoOrder)
	return err
}

func (r *OrderRepositoryMongo) FindOrderByID(orderId string) (*entities.Order, error) {
	var mOrder mongoOrder
	filter := bson.M{"_id": orderId}
	collection := r.db.Collection(orderCollectionName)
	err := collection.FindOne(context.TODO(), filter).Decode(&mOrder)
	if err != nil {
		return nil, err
	}
	return toOrder(&mOrder), nil
}

func (r *OrderRepositoryMongo) Update(order *entities.Order) error {
	mOrder := toMongoOrder(order)
	filter := bson.M{"_id": order.GetId()}
	update := bson.M{"$set": bson.M{
		"customer_id": mOrder.ID,
		"items": mOrder.Items,
		"payment_status": mOrder.PaymentStatus,
		"preparation_status": mOrder.PreparationStatus,
	}}
	collection := r.db.Collection(orderCollectionName)
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}
