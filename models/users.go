package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name      string             `json:"name" bson:"name"`
    Email     string             `json:"email" bson:"email"`
    Password  string             `json:"password,omitempty" bson:"password,omitempty"`
    Role      string             `json:"role" bson:"role"`
    Phone     string             `json:"phone" bson:"phone"`
    Address   string             `json:"address" bson:"address"`
    IsActive  bool               `json:"isActive" bson:"isActive"`
    CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
    UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

