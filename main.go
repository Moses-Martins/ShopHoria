package main

import(
	"net/http"
	"github.com/gorilla/mux"
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
	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
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

	clientOptions := options.Client().ApplyURI(dbURL)
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

	DB = client.Database(dbURL)		

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

	router := mux.NewRouter()
	router.HandleFunc("/api/auth/register", apiCfg.handlerRegister).Methods("POST")



	
	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}