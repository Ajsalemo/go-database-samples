package getdata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	c "github.com/Ajsalemo/go-database-samples/cosmosdb-mongodb/config/db"
)

func ApiRoute(w http.ResponseWriter, r *http.Request) {
	type Todo struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Name      string             `bson:"name"`
		Completed bool               `bson:"completed"`
	}
	var todos []Todo
	todoTable := [][]string{}

	database := os.Getenv("COSMOS_DB_DATABASE_NAME")
	collection := os.Getenv("COSMOS_DB_COLLECTION_NAME")
	// Connect to the database
	c := c.Connect()
	ctx := context.Background()

	defer c.Disconnect(ctx)

	// Make a reference to our collection
	t := c.Database(database).Collection(collection)
	// Find all documents
	rows, e := t.Find(ctx, bson.D{})
	if e != nil {
		log.Fatalf("failed to list todo(s) %v", e)
	}

	e = rows.All(ctx, &todos)

	if e != nil {
		log.Fatalf("failed to list todo(s) %v", e)
	}

	for _, todo := range todos {
		s, _ := todo.ID.MarshalJSON()
		b := strconv.FormatBool(todo.Completed)
		todoTable = append(todoTable, []string{string(s), todo.Name, b})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(todoTable)
	json.NewEncoder(w).Encode(todoTable)
}
