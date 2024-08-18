package services

import (
    "database/sql"
    "github.com/GarvitDadheech/quiz-app-backend/models"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

var db *sql.DB

// Initialize the services with the database connection
func Initialize(database *sql.DB) {
    db = database
}

func LoginUser(username, password string) (int, error) {
    var dbUser models.User
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, err
        }
        return 0, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
    if err != nil {
        return 0, err
    }

    return dbUser.ID, nil
}

func RegisterUser(username, password string) (int64, string, error) {
    // Check if the username already exists
    var existingUserId int64
    var existingUserPassword string
    err := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&existingUserId, &existingUserPassword)
    if err != nil && err != sql.ErrNoRows {
        return 0, "", err
    }

    if existingUserId != 0 {
        // Check if the password matches
        if bcrypt.CompareHashAndPassword([]byte(existingUserPassword), []byte(password)) == nil {
            // Username and password match, user already exists
            return 0, "User already exists. Please login.", nil
        } else {
            // Username exists but password doesn't match
            return 0, "Username already taken. Please choose a different username.", nil
        }
    }

    // Proceed with registration
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return 0, "", err
    }

    result, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedPassword))
    if err != nil {
        return 0, "", err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, "", err
    }

    
    _, err = db.Exec(`
        INSERT INTO user_badges (user_id, badge_id, earned)
        SELECT ?, b.id, FALSE
        FROM badges b
    `, id)
    if err != nil {
        return 0, "", err
    }

    return id, "", nil
}



func FetchQuizzes() ([]models.Quiz, error) {
    rows, err := db.Query("SELECT id, title, description FROM quizzes")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var quizzes []models.Quiz
    for rows.Next() {
        var quiz models.Quiz
        if err := rows.Scan(&quiz.ID, &quiz.Title, &quiz.Description); err != nil {
            return nil, err
        }
        quizzes = append(quizzes, quiz)
    }

    return quizzes, nil
}

func FetchQuiz(id string) (models.QuizWithQuestions, error) {
    var quiz models.QuizWithQuestions
    err := db.QueryRow("SELECT title, description FROM quizzes WHERE id = ?", id).Scan(&quiz.Title, &quiz.Description)
    if err != nil {
        return quiz, err
    }

    rows, err := db.Query("SELECT id, question FROM questions WHERE quiz_id = ?", id)
    if err != nil {
        return quiz, err
    }
    defer rows.Close()

    var questions []models.QuestionWithAnswers
    for rows.Next() {
        var question models.QuestionWithAnswers
        if err := rows.Scan(&question.ID, &question.Question); err != nil {
            return quiz, err
        }

        aRows, err := db.Query("SELECT id, answer, is_correct FROM answers WHERE question_id = ?", question.ID)
        if err != nil {
            return quiz, err
        }
        defer aRows.Close()

        var answers []models.Answer
        for aRows.Next() {
            var answer models.Answer
            if err := aRows.Scan(&answer.ID, &answer.Answer, &answer.IsCorrect); err != nil {
                return quiz, err
            }
            answers = append(answers, answer)
        }

        question.Answers = answers
        questions = append(questions, question)
    }

    quiz.Questions = questions
    return quiz, nil
}

func SubmitAnswer(userId, quizId, questionId, answerId int) (map[string]interface{}, error) {
    var isCorrect bool
    err := db.QueryRow("SELECT is_correct FROM answers WHERE id = ? AND question_id = ?", answerId, questionId).Scan(&isCorrect)
    if err != nil {
        return nil, err
    }

    var score int
    if isCorrect {
        score = 1
    }
    _, err = db.Exec("INSERT INTO user_scores (user_id, quiz_id, score) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE score = score + ?", userId, quizId, score, score)
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "message": "Answer submitted",
        "correct": isCorrect,
    }, nil
}

func FetchUserScore(userId string) ([]models.UserScore, error) {
    rows, err := db.Query("SELECT quiz_id, score FROM user_scores WHERE user_id = ?", userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var scores []models.UserScore
    for rows.Next() {
        var score models.UserScore
        if err := rows.Scan(&score.QuizID, &score.Score); err != nil {
            return nil, err
        }
        scores = append(scores, score)
    }

    return scores, nil
}

func UpdateUserCash(userID int, cashAmount float64) error {
    _, err := db.Exec("UPDATE users SET cash = ? WHERE id = ?", cashAmount, userID)
    if err != nil {
        return err
    }
    return nil
}

func FetchUserCash(userId string) (float64, error) {
    var cashAmount float64
    err := db.QueryRow("SELECT cash FROM users WHERE id = ?", userId).Scan(&cashAmount)
    if err != nil {
        return 0, err
    }
    return cashAmount, nil
}

func GetLeaderboard() ([]map[string]interface{}, error) {
    rows, err := db.Query("SELECT u.id, u.username, u.cash from users u ORDER BY u.cash DESC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var leaderboard []map[string]interface{}
    for rows.Next() {
        var userID int
        var username string
        var cash float64
        if err := rows.Scan(&userID, &username, &cash); err != nil {
            return nil, err
        }
        leaderboard = append(leaderboard, map[string]interface{}{
            "user_id": userID,
            "username": username,
            "cash":    cash,
        })
    }
    return leaderboard, nil
}

func GetRecentQuizID(userID string) (int, error) {
    var recentQuizID int
    err := db.QueryRow("SELECT recent_quiz_id FROM users WHERE id = ?", userID).Scan(&recentQuizID)
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, nil
        }
        return 0, err
    }
    return recentQuizID, nil
}

func UpdateRecentQuiz(userID int, quizID int) error {
    _, err := db.Exec("UPDATE users SET recent_quiz_id = ? WHERE id = ?", quizID, userID)
    if err != nil {
        return err
    }
    return nil
}

func DeleteUserById(userId int) error {
    // Delete from user_badges table
    _, err := db.Exec("DELETE FROM user_badges WHERE user_id = ?", userId)
    if err != nil {
        return err
    }

    // Delete from user_scores table
    _, err = db.Exec("DELETE FROM user_scores WHERE user_id = ?", userId)
    if err != nil {
        return err
    }

    // Delete from user_quizzes table
    _, err = db.Exec("DELETE FROM user_quizzes WHERE user_id = ?", userId)
    if err != nil {
        return err
    }

    // Delete from users table
    _, err = db.Exec("DELETE FROM users WHERE id = ?", userId)
    if err != nil {
        return err
    }

    return nil
}



type Badge struct {
    BadgeID     int    `json:"badge_id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Earned      bool   `json:"earned"`
}

// FetchAllUserBadges retrieves all badges for a given user
func FetchAllUserBadges(userID string) ([]Badge, error) {
    var badges []Badge

    query := `
        SELECT b.id AS badge_id, b.name, b.description, COALESCE(ub.earned, FALSE) AS earned
        FROM badges b
        LEFT JOIN user_badges ub ON b.id = ub.badge_id AND ub.user_id = ?
    `
    
    rows, err := db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var badge Badge
        if err := rows.Scan(&badge.BadgeID, &badge.Name, &badge.Description, &badge.Earned); err != nil {
            return nil, err
        }
        badges = append(badges, badge)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return badges, nil
}

func GetUsernameByID(userID string) (string, error) {
    var username string
    err := db.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
    if err == sql.ErrNoRows {
        return "", nil
    } else if err != nil {
        return "", err
    }
    return username, nil
}

func UpdateBadges(userID int, badgeIDs []int) error {
    // Iterate over badgeIDs and perform the update for each
    for _, badgeID := range badgeIDs {
        query := "UPDATE user_badges SET earned = TRUE WHERE user_id = ? AND badge_id = ?"
        _, err := db.Exec(query, userID, badgeID)
        if err != nil {
            return fmt.Errorf("could not execute statement for badge ID %d: %w", badgeID, err)
        }
    }
    
    return nil
}

func FetchBadgeNames(badgeIDs []int) (map[int]string, error) {
    badgeNames := make(map[int]string)

    for _, badgeID := range badgeIDs {
        query := "SELECT name FROM badges WHERE id = ?"
        var badgeName string
        err := db.QueryRow(query, badgeID).Scan(&badgeName)
        if err != nil {
            // Log the error for debugging
            fmt.Printf("Error retrieving name for badge ID %d: %v\n", badgeID, err)
            return nil, fmt.Errorf("could not retrieve name for badge ID %d: %w", badgeID, err)
        }
        badgeNames[badgeID] = badgeName
    }

    return badgeNames, nil
}

