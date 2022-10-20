package getdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/Ajsalemo/go-database-samples/mysql/config/db"
)

// NOTE: The properties MUST be uppercase, even if the fields we're trying to access in MySQL are lowercase
// i.e todos, completed, id
// Or else json encoding won't succeed and we'll only be able to print this to stream. While the browser just shows an empty object (s)
func ApiRoute(w http.ResponseWriter, r *http.Request) {
	type Todo struct {
		Id   string  	`json:"id"`
		Todos string 	`json:"name"`
		Completed bool 	`json:"completed"`
	}

	todoResponse := []Todo{}

	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos")

	if err != nil {
        log.Print(err)
    }

	for rows.Next() {
		var todo Todo
        err = rows.Scan(&todo.Id, &todo.Todos, &todo.Completed)
        if err != nil {
            log.Fatal(err.Error())
        } 
        todoResponse = append(todoResponse, todo)
    }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println(todoResponse)
	json.NewEncoder(w).Encode(todoResponse)
}
