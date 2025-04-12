package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		OrderEntry: OrderEntry{},
	}
}

type Models struct {
	OrderEntry OrderEntry
}

type OrderEntry struct {
	ID         string    `json:"id,omitempty" bson:"_id,omitempty"`
	OrderID    string    `json:"orderId" bson:"order_id"`
	CustomerID string    `json:"customerId" bson:"customer_id"`
	OrderDate  time.Time `json:"orderDate" bson:"order_date"`
	Status     string    `json:"status" bson:"status"`
	Subtotal   float64   `json:"subtotal" bson:"subtotal"`
	Tax        float64   `json:"tax" bson:"tax"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func (o *OrderEntry) Insert(entry OrderEntry) error {
	collection := client.Database("orders").Collection("orders")

	_, err := collection.InsertOne(context.TODO(), OrderEntry{
		OrderID:    entry.OrderID,
		CustomerID: entry.CustomerID,
		OrderDate:  entry.OrderDate,
		Status:     entry.Status,
		//Items:           entry.Items,
		Subtotal: entry.Subtotal,
		Tax:      entry.Tax,
		//ShippingCost:    entry.ShippingCost,
		//Total:           entry.Total,
		//ShippingMethod:  entry.ShippingMethod,
		//PaymentMethod:   entry.PaymentMethod,
		//BillingAddress:  entry.BillingAddress,
		//ShippingAddress: entry.ShippingAddress,
		//Notes:           entry.Notes,
		//Discounts:       entry.Discounts,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into orders:", err)
		return err
	}

	return nil
}
