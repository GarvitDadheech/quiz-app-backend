package main

import (
    "github.com/gin-gonic/gin"

    "github.com/GarvitDadheech/quiz-app-backend/controllers"
    "github.com/GarvitDadheech/quiz-app-backend/services"
    "github.com/GarvitDadheech/quiz-app-backend/database"
)

func main() {
    database.InitDB()
    defer database.GetDB().Close()

    services.Initialize(database.GetDB())

    r := gin.Default()
    r.POST("/login", controllers.Login)
    r.POST("/register", controllers.Register)
    r.GET("/quiz/:id", controllers.GetQuiz)
    r.POST("/submit-answer", controllers.SubmitAnswer)
    r.GET("/user-cash/:userId", controllers.GetUserCash)
    r.POST("/update-user-cash", controllers.UpdateUserCash)
    r.GET("/leaderboard", controllers.GetLeaderboard)
    r.GET("/recent-quiz/:user_id", controllers.GetRecentQuiz)
    r.PUT("/update-recent-quiz", controllers.UpdateRecentQuiz)
    r.DELETE("/delete-user/:userId", controllers.DeleteUser)
    r.GET("/user/:userId/badges", controllers.GetAllUserBadges)
    r.GET("/user/:userId/username", controllers.GetUsernameHandler)
    r.POST("/update-user-badges", controllers.UpdateUserBadges)
    r.POST("/get-badge-names", controllers.GetBadgeNames)
    r.Run(":8080")
}