package models


type Category struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`        
    Name        string             `json:"name" bson:"name"`                         
    Description string             `json:"description,omitempty" bson:"description,omitempty"`
    CreatedAt   time.Time          `json:"created_at" bson:"created_at"`           
}
