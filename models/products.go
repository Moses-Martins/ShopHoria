package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
    ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
    VendorID    primitive.ObjectID   `json:"vendor_id" bson:"vendor_id"`        
    CategoryID  primitive.ObjectID   `json:"category_id" bson:"category_id"`    
    Name        string               `json:"name" bson:"name"`
    Description string               `json:"description" bson:"description"`
    Price       float64              `json:"price" bson:"price"`                 
    Stock       int                  `json:"stock" bson:"stock"`
    Images      []string             `json:"images" bson:"images"`
    IsActive    bool                 `json:"is_active" bson:"is_active"`
    CreatedAt   time.Time            `json:"created_at" bson:"created_at"`
    UpdatedAt   time.Time            `json:"updated_at" bson:"updated_at"`
}

