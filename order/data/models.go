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
	ID              string      `json:"id,omitempty" bson:"_id,omitempty"`
	OrderID         string      `json:"orderId" bson:"order_id"`
	CustomerID      string      `json:"customerId" bson:"customer_id"`
	OrderDate       time.Time   `json:"orderDate" bson:"order_date"`
	Status          string      `json:"status" bson:"status"`
	Items           []OrderItem `json:"items" bson:"items"`
	Subtotal        float64     `json:"subtotal" bson:"subtotal"`
	Tax             float64     `json:"tax" bson:"tax"`
	ShippingCost    float64     `json:"shippingCost" bson:"shipping_cost"`
	Total           float64     `json:"total" bson:"total"`
	ShippingMethod  string      `json:"shippingMethod" bson:"shipping_method"`
	PaymentMethod   string      `json:"paymentMethod" bson:"payment_method"`
	BillingAddress  Address     `json:"billingAddress" bson:"billing_address"`
	ShippingAddress Address     `json:"shippingAddress" bson:"shipping_address"`
	Notes           string      `json:"notes,omitempty" bson:"notes,omitempty"`
	Discounts       []Discount  `json:"discounts,omitempty" bson:"discounts,omitempty"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type OrderItem struct {
	ProductID   string  `json:"productId" bson:"product_id"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	UnitPrice   float64 `json:"unitPrice" bson:"unit_price"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	SKU         string  `json:"sku" bson:"sku"`
}

type Address struct {
	Street     string `json:"street" bson:"street"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	PostalCode string `json:"postalCode" bson:"postal_code"`
	Country    string `json:"country" bson:"country"`
}

type Discount struct {
	Code        string  `json:"code" bson:"code"`
	Amount      float64 `json:"amount" bson:"amount"`
	IsPercent   bool    `json:"isPercent" bson:"is_percent"`
	Description string  `json:"description" bson:"description"`
}

func (o *OrderEntry) Insert(entry OrderEntry) error {
	collection := client.Database("orders").Collection("orders")

	_, err := collection.InsertOne(context.TODO(), OrderEntry{
		OrderID:         entry.OrderID,
		CustomerID:      entry.CustomerID,
		OrderDate:       entry.OrderDate,
		Status:          entry.Status,
		Items:           entry.Items,
		Subtotal:        entry.Subtotal,
		Tax:             entry.Tax,
		ShippingCost:    entry.ShippingCost,
		Total:           entry.Total,
		ShippingMethod:  entry.ShippingMethod,
		PaymentMethod:   entry.PaymentMethod,
		BillingAddress:  entry.BillingAddress,
		ShippingAddress: entry.ShippingAddress,
		Notes:           entry.Notes,
		Discounts:       entry.Discounts,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into orders:", err)
		return err
	}

	return nil
}
