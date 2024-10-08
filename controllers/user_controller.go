package controllers

import (
    "database/sql"
    "net/http"
    "strconv"
    "github.com/GarvitDadheech/quiz-app-backend/services"
    "github.com/GarvitDadheech/quiz-app-backend/models"
    "github.com/gin-gonic/gin"
    "fmt"
)

// Login handles the login request
func Login(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId, err := services.LoginUser(user.Username, user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": userId})
}

// Register handles the user registration request
func Register(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId, message, err := services.RegisterUser(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if message != "" {
        if message == "user already exists, kindly login" {
            c.JSON(http.StatusConflict, gin.H{"error": message})
        } else {
            c.JSON(http.StatusConflict, gin.H{"error": message})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": userId})
}




func GetQuizzes(c *gin.Context) {
    quizzes, err := services.FetchQuizzes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quizzes)
}

func GetQuiz(c *gin.Context) {
    id := c.Param("id")
    quiz, err := services.FetchQuiz(id)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, quiz)
}


func SubmitAnswer(c *gin.Context) {
    var submission struct {
        UserId     int `json:"user_id"`
        QuizId     int `json:"quiz_id"`
        QuestionId int `json:"question_id"`
        AnswerId   int `json:"answer_id"`
    }

    if err := c.BindJSON(&submission); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := services.SubmitAnswer(submission.UserId, submission.QuizId, submission.QuestionId, submission.AnswerId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}


func GetUserScore(c *gin.Context) {
    userId := c.Param("userId")
    scores, err := services.FetchUserScore(userId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user_id": userId,
        "scores":  scores,
    })
}

func GetUserCash(c *gin.Context) {
    userId := c.Param("userId")
    cashAmount, err := services.FetchUserCash(userId)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "User cash not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, gin.H{"cash_amount": cashAmount})
}

func UpdateUserCash(c *gin.Context) {
    var updateData struct {
        UserID    int     `json:"user_id"`
        CashAmount float64 `json:"cash_amount"`
    }

    if err := c.BindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.UpdateUserCash(updateData.UserID, updateData.CashAmount)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User cash updated successfully"})
}

func GetLeaderboard(c *gin.Context) {
    leaderboard, err := services.GetLeaderboard()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve leaderboard"})
        return
    }
    c.JSON(http.StatusOK, leaderboard)
}

func GetRecentQuiz(c *gin.Context) {
    userID := c.Param("user_id")

    recentQuizID, err := services.GetRecentQuizID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, recentQuizID)
}

func UpdateRecentQuiz(c *gin.Context) {
    var updateData struct {
        UserID int `json:"user_id"`
        QuizID int `json:"quiz_id"`
    }

    if err := c.BindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.UpdateRecentQuiz(updateData.UserID, updateData.QuizID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Recent quiz updated successfully"})
}

func DeleteUser(c *gin.Context) {
    userIdStr := c.Param("userId")
    userId, err := strconv.Atoi(userIdStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to convert userId"})
        return
    }

    err = services.DeleteUserById(userId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func GetAllUserBadges(c *gin.Context) {
    userID := c.Param("userId")

    badges, err := services.FetchAllUserBadges(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch badges"})
        return
    }
    
    badgeMap := make(map[string]map[string]interface{})
    for _, badge := range badges {
        badgeMap[badge.Name] = map[string]interface{}{
            "description": badge.Description,
            "earned":      badge.Earned,
        }
    }

    c.JSON(http.StatusOK, badgeMap)
}

func GetUsernameHandler(c *gin.Context) {
    userID := c.Param("userId")
    
    username, err := services.GetUsernameByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        return
    }
    
    if username == "" {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"username": username})
}

func UpdateUserBadges(c *gin.Context) {
    var requestBody struct {
        UserID    int   `json:"user_id"`
        BadgeIDs  []int `json:"badge_ids"`
    }

   
    if err := c.BindJSON(&requestBody); err != nil {
      
        fmt.Printf("Failed to bind JSON: %v\n", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    userID := requestBody.UserID
    badgeIDs := requestBody.BadgeIDs

 
    fmt.Printf("Received request - UserID: %d, BadgeIDs: %v\n", userID, badgeIDs)

    
    err := services.UpdateBadges(userID, badgeIDs)
    if err != nil {
       
        fmt.Printf("Failed to update badges: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update badges"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Badges updated successfully"})
}

func GetBadgeNames(c *gin.Context) {
    var requestBody struct {
        BadgeIDs []int `json:"badge_ids"`
    }
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    badgeIDs := requestBody.BadgeIDs
    badgeNames, err := services.FetchBadgeNames(badgeIDs)
    if err != nil {
        fmt.Printf("Failed to fetch badge names: %v\n", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch badge names"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"badge_names": badgeNames})
}

