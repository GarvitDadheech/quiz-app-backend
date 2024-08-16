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
    dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database connection: ", err)
    }
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    log.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
    return db
}
