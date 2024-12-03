package queries

import (
	"context"
	"tech-challenge-fase-1/internal/core/dtos"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type OrderDisplayListQueryMongo struct {
	db *mongo.Database
}

func NewOrderDisplayListQueryMongo(db *mongo.Database) *OrderDisplayListQueryMongo {
	return &OrderDisplayListQueryMongo{db: db}
}

const orderCollectionName = "orders"

func (q *OrderDisplayListQueryMongo) Execute() ([]*dtos.OrderDisplayDTO, error) {
	collection := q.db.Collection(orderCollectionName)
	filter := bson.M{
		"preparation_status": bson.M{
			"$nin": []string{"AWAITING","CANCELED","FINISHED"},
		},
	}
	opts := options.Find()
	opts.SetSort(bson.D{{"preparation_status", 1}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var orders []*dtos.OrderDisplayDTO
	if err := cursor.All(context.TODO(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}
