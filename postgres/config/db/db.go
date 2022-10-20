package db

import (
    "database/sql"
    "os"
    "fmt"
)
 
func Connect() *sql.DB {
    dbDriver := "postgres"
    postgresUser := os.Getenv("POSTGRES_USER")
    postgresPassword := os.Getenv("POSTGRES_PASSWORD")
    postgresName := os.Getenv("POSTGRES_DATABASE_NAME")
    postgresHost := os.Getenv("POSTGRES_DATABASE_HOST")

    formattedConnStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", postgresHost, postgresUser, postgresPassword, postgresName)
    db, err := sql.Open(dbDriver, formattedConnStr)
    if err != nil {
        panic(err.Error())
    }
    return db
}