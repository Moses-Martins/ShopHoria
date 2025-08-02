package models

type Cart struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`       // Primary Key
    UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`                 // FK → Users.id (buyer)
    ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`           // FK → Products.id
    Quantity  int                `json:"quantity" bson:"quantity"`               // Number of items
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`           // Auto-generated
}