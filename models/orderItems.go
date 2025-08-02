package models

type OrderItem struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`        // Primary Key
    OrderID   primitive.ObjectID `json:"order_id" bson:"order_id"`                // FK → Orders.id
    ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`            // FK → Products.id
    Quantity  int                `json:"quantity" bson:"quantity"`                // Number of items
    Price     float64            `json:"price" bson:"price"`                      // Price at time of purchase
}