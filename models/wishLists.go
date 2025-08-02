package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wishlist struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`     // Primary Key
    UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`               // FK → Users.id (buyer)
    ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`         // FK → Products.id
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`         // Auto-generated
}
