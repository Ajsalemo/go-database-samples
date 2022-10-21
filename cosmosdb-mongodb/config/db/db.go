package db

import (
    "context"
    "log"
    "time"
    "os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
 
func Connect() *mongo.Client {
    mongoDbConnStr := os.Getenv("COSMOS_DB_CONN_STR")
    
    // If we don't set these timeout options / cancel then we'll get a 'client has disconnected' immediately after connecting
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

    clientOptions := options.Client().ApplyURI(mongoDbConnStr).SetDirect(true)
    c, err := mongo.NewClient(clientOptions)

    err = c.Connect(ctx)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

    return c
}