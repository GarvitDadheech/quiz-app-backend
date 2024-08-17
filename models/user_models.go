package models

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Cash     float64 `json:"cash"` 
}

type Quiz struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type QuestionWithAnswers struct {
    ID       int      `json:"id"`
    Question string   `json:"question"`
    Answers  []Answer `json:"answers"`
}

type Answer struct {
    ID         int  `json:"id"`
    Answer     string `json:"answer"`
    IsCorrect  bool   `json:"is_correct"`
}

type QuizWithQuestions struct {
    Title       string                `json:"title"`
    Description string                `json:"description"`
    Questions   []QuestionWithAnswers `json:"questions"`
}

type UserScore struct {
    QuizID int `json:"quiz_id"`
    Score  int `json:"score"`
}
