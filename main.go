package main

import(
	"os"
	"log"
	"net/http"
	"context"
	"time"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv" 
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type apiConfig struct {
	port string
	JwtSecret string
	DB *mongo.Database
	RegisterRedirectUrl string
	LoginRedirectUrl string
	GoogleOauthConfig *oauth2.Config
	assetsRoot       string
}


func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	jwtSecret := os.Getenv("SECRET")
	assetsRoot := os.Getenv("ASSETS_ROOT")
	RegisterRedirectUrl :=  os.Getenv("REGISTER_REDIRECT_URL")
	LoginRedirectUrl :=  os.Getenv("LOGIN_REDIRECT_URL")
	clientID :=     os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	scopes := []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		}
	endpoint := google.Endpoint

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB := client.Database(mongoDB)		

	apiCfg := apiConfig {
		port: port,
		DB: DB,
		JwtSecret: jwtSecret,
		assetsRoot: assetsRoot,
		RegisterRedirectUrl: RegisterRedirectUrl,
		LoginRedirectUrl: LoginRedirectUrl,
		GoogleOauthConfig: &oauth2.Config{
			ClientID:    clientID,
			ClientSecret: clientSecret,
			Scopes: scopes,
			Endpoint: endpoint,
		},
	}


	apiCfg.InitCollections()
	router := mux.NewRouter()
	router.HandleFunc("/api/auth/register", apiCfg.handlerRegister).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}



func (cfg *apiConfig) InitCollections() {
	Users = cfg.DB.Collection("users")
	/*Carts = config.DB.Collection("carts")
	OrderItems = config.DB.Collection("order_items")
	Payments = config.DB.Collection("payments")
	Products = config.DB.Collection("products")
	WishList = config.DB.Collection("wish_list")
	Orders = config.DB.Collection("orders")
	Categories = config.DB.Collection("categories")*/
}