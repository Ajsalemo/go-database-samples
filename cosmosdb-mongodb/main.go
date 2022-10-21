package main

import (
	"log"
	"net/http"

	index "github.com/Ajsalemo/go-database-samples/cosmosdb-mongodb/routes/index"
	getdata "github.com/Ajsalemo/go-database-samples/cosmosdb-mongodb/routes/getdata"
)

func main() {
	http.HandleFunc("/", index.IndexRoute)
	http.HandleFunc("/api/get", getdata.ApiRoute)

	log.Printf("server listening at 0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Failed to start the HTTP server!")
	}
}