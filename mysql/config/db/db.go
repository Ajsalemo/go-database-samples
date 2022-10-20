package db

import (
    "database/sql"
    "os"
)
 
func Connect() *sql.DB {
    dbDriver := "mysql"
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    dbName := os.Getenv("MYSQL_DATABASE_NAME")
    dbHost := os.Getenv("MYSQL_DATABASE_HOST")
 
    db, err := sql.Open(dbDriver, mysqlUser + ":" + mysqlPassword + "@" + "tcp(" + dbHost + ")" + "/" + dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}