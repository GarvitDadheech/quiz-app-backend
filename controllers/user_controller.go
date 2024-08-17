package controllers

import (
    "database/sql" // Add this import
    "net/http"

    "github.com/GarvitDadheech/quiz-app-backend/services"
    "github.com/GarvitDadheech/quiz-app-backend/models"
    "github.com/gin-gonic/gin"
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

    userId, err := services.RegisterUser(user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": userId})
}

// GetQuizzes retrieves all quizzes
func GetQuizzes(c *gin.Context) {
    quizzes, err := services.FetchQuizzes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, quizzes)
}

// GetQuiz retrieves a specific quiz by its ID
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

// SubmitAnswer handles answer submissions
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

// GetUserScore retrieves the score of a user
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
