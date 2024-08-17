package main

import (
    "database/sql"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"

    "github.com/GarvitDadheech/quiz-app-backend/controllers"
    "github.com/GarvitDadheech/quiz-app-backend/services"
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

func main() {
    initDB()

    // Initialize services with the database connection
    services.Initialize(db)

    r := gin.Default()
    r.POST("/login", controllers.Login)
    r.POST("/register", controllers.Register)
    r.GET("/quizzes", controllers.GetQuizzes)
    r.GET("/quiz/:id", controllers.GetQuiz)
    r.POST("/submit-answer", controllers.SubmitAnswer)
    r.GET("/user-score/:userId", controllers.GetUserScore)
    r.GET("/user-cash/:userId", controllers.GetUserCash)
    r.POST("/update-user-cash", controllers.UpdateUserCash)
    r.GET("/leaderboard", controllers.GetLeaderboard)
    r.GET("/recent-quiz/:user_id", controllers.GetRecentQuiz)
    r.Run(":8080")
}
