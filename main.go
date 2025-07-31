package main

import(
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Serving on: http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}