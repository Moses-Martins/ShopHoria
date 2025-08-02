package main

import(
	"net/http"
	"log"
	"time"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/Moses-Martins/ShopHoria/models"
	"github.com/Moses-Martins/ShopHoria/internal/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type regInput struct {
	Password string `json:"password"`
	Email string `json:"email"`
	Name string `json:"name"`
}

type returnUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
    Name      string             `json:"name" bson:"name"`
    Email     string             `json:"email" bson:"email"`
    Role      string             `json:"role" bson:"role"`
    IsActive  bool               `json:"isActive" bson:"isActive"`
    CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
    UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

var Users *mongo.Collection


func (cfg *apiConfig) handlerRegister(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	params := regInput{}


	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	
	count, _ := Users.CountDocuments(req.Context(), bson.M{"email": params.Email})
	if count > 0 {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	params.Password, err = auth.HashPassword(params.Password)
	if err != nil {
		log.Printf("Error Hashing Password: %s", err)
		w.WriteHeader(500)
		return
	}

	User := models.User{
		Name: params.Name,
		Email: params.Email,
		Password: params.Password,
		Role: "buyer",
		IsActive: true,
		CreatedAt: time.Now(), 
		UpdatedAt: time.Now(),
	}

	result, err := Users.InsertOne(req.Context(), User)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	objID := result.InsertedID.(primitive.ObjectID)

	var insertedUser models.User
	err = Users.FindOne(req.Context(), bson.M{"_id": objID}).Decode(&insertedUser)
	if err != nil {
		http.Error(w, "Failed to fetch inserted user", http.StatusInternalServerError)
		return
	}

	respBody := returnUser{
		ID: insertedUser.ID,
		Name: insertedUser.Name,
		Email: insertedUser.Email,
		Role: insertedUser.Role,
		IsActive: insertedUser.IsActive,
		CreatedAt: insertedUser.CreatedAt,
		UpdatedAt: insertedUser.UpdatedAt,
	}

	data, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(data)
	
}
