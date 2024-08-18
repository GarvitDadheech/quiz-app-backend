package database

import (
    "database/sql"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    
    dbUsername := "root"
    dbPassword := "mypassword"
    dbHost := "localhost"
    dbPort := "3306"
    dbName := "quiz_app"

    dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database connection: ", err)
    }

    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(25)
    DB.SetConnMaxLifetime(5 * time.Minute)

    err = DB.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }

    log.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
    return DB
}