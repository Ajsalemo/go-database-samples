package db

import (
    "database/sql"
    "os"
)
 
func Connect() *sql.DB {
    dbDriver := "mysql"
    mariaDbUser := os.Getenv("MARIADB_USER")
    mariaDbPassword := os.Getenv("MARIADB_PASSWORD")
    mariaDatabaseName := os.Getenv("MARIADB_DATABASE_NAME")
    mariaDbHost := os.Getenv("MARIADB_DATABASE_HOST")
 
    db, err := sql.Open(dbDriver, mariaDbUser + ":" + mariaDbPassword + "@" + "tcp(" + mariaDbHost + ")" + "/" + mariaDatabaseName)
    if err != nil {
        panic(err.Error())
    }
    return db
}