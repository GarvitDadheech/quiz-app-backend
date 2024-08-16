package main

import (
    "database/sql"
    "log"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
    var err error
    
    dbUsername := "root"
    dbPassword := "mypassword"
    dbHost := "localhost"
    dbPort := "3306"
    dbName := "quiz_app"

    // Create the data source name (DSN)
    dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

    // Open a connection to the database
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database connection: ", err)
    }

    // Set connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Verify the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }

    log.Println("Successfully connected to the database")
}
