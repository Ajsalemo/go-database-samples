package db

import (
    "database/sql"
    "os"
)
 
func Connect() *sql.DB {
    dbDriver := "sqlserver"
    mssqlUser := os.Getenv("MSSQL_USER")
    mssqlPassword := os.Getenv("MSSQL_PASSWORD")
    mssqlDatabaseName := os.Getenv("MSSQL_DATABASE_NAME")
    mssqlHost := os.Getenv("MSSQL_DATABASE_HOST")

    mssqlConnStr := dbDriver + "://" + mssqlUser + ":" + mssqlPassword + "@" + mssqlHost + "?database=" + mssqlDatabaseName
    db, err := sql.Open(dbDriver, mssqlConnStr)
    if err != nil {
        panic(err.Error())
    }
    return db
}