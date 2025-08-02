package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
    ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`               // Primary Key
    UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`                         // FK → Users.id
    OrderID        primitive.ObjectID `json:"order_id" bson:"order_id"`                       // FK → Orders.id
    Amount         float64            `json:"amount" bson:"amount"`                           // Decimal(10,2)
    Status         string      `json:"status" bson:"status"`                           // Enum: pending, completed, failed
    PaymentMethod  string      `json:"payment_method" bson:"payment_method"`           // Enum: stripe, paypal, etc.
    TransactionRef string             `json:"transaction_ref" bson:"transaction_ref"`         // Gateway transaction reference
    CreatedAt      time.Time          `json:"created_at" bson:"created_at"`                   // Auto-generated
}