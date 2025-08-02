package models

type Order struct {
    ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`                  // Primary Key
    BuyerID         primitive.ObjectID `json:"buyer_id" bson:"buyer_id"`                          // FK → Users.id (buyer)
    VendorID        primitive.ObjectID `json:"vendor_id" bson:"vendor_id"`                        // FK → Users.id (seller)
    Status          string        	   `json:"status" bson:"status"`                              // Enum: pending, paid, etc.
    TotalAmount     float64            `json:"total_amount" bson:"total_amount"`                  // Decimal(10,2)
    PaymentID       primitive.ObjectID `json:"payment_id" bson:"payment_id"`                      // FK → Payments.id
    ShippingAddress string             `json:"shipping_address" bson:"shipping_address"`          // Delivery address
    CreatedAt       time.Time          `json:"created_at" bson:"created_at"`                      // Auto-generated
    UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`                      // Auto-generated
}
