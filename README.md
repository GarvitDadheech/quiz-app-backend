# Quiz App Backend API Documentation

Welcome to the documentation for the Quiz App backend API. This API is built using Go (Golang) with the Gin framework and provides various endpoints for user management, quiz operations, and more.

## Base URL

All API endpoints are available at the base URL:

```
http://localhost:8080
```

## Endpoints

### User Management

#### 1. **Login**

**`POST /login`**

- **Description:** Authenticates a user and returns a user ID upon successful login.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "Login successful",
      "user_id": "number"
    }
    ```
  - **400 Bad Request** or **401 Unauthorized**
    ```json
    {
      "error": "error message"
    }
    ```

#### 2. **Register**

**`POST /register`**

- **Description:** Registers a new user. If the user already exists, an appropriate message is returned.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "User registered successfully",
      "user_id": "number"
    }
    ```
  - **409 Conflict**
    ```json
    {
      "error": "error message"
    }
    ```

#### 3. **Delete User**

**`DELETE /delete-user/:userId`**

- **Description:** Deletes a user by their ID.
- **Path Parameter:**
  - `userId` - The ID of the user to delete.
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "User deleted successfully"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 4. **Get Username**

**`GET /user/:userId/username`**

- **Description:** Retrieves the username for a given user ID.
- **Path Parameter:**
  - `userId` - The ID of the user.
- **Responses:**
  - **200 OK**
    ```json
    {
      "username": "string"
    }
    ```
  - **404 Not Found** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

### Quiz Management

#### 5. **Get Quiz**

**`GET /quiz/:id`**

- **Description:** Retrieves a specific quiz by its ID.
- **Path Parameter:**
  - `id` - The ID of the quiz.
- **Responses:**
  - **200 OK**
    ```json
    {
      "quiz": "quiz data"
    }
    ```
  - **404 Not Found** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 6. **Submit Answer**

**`POST /submit-answer`**

- **Description:** Submits an answer for a specific quiz question.
- **Request Body:**
  ```json
  {
    "user_id": "number",
    "quiz_id": "number",
    "question_id": "number",
    "answer_id": "number"
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "result": "result data"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 7. **Get Recent Quiz**

**`GET /recent-quiz/:user_id`**

- **Description:** Retrieves the ID of the most recent quiz taken by a user.
- **Path Parameter:**
  - `user_id` - The ID of the user.
- **Responses:**
  - **200 OK**
    ```json
    {
      "quiz_id": "number"
    }
    ```
  - **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 8. **Update Recent Quiz**

**`PUT /update-recent-quiz`**

- **Description:** Updates the most recent quiz ID for a user.
- **Request Body:**
  ```json
  {
    "user_id": "number",
    "quiz_id": "number"
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "Recent quiz updated successfully"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

### User Cash Management

#### 9. **Get User Cash**

**`GET /user-cash/:userId`**

- **Description:** Retrieves the cash amount for a given user ID.
- **Path Parameter:**
  - `userId` - The ID of the user.
- **Responses:**
  - **200 OK**
    ```json
    {
      "cash_amount": "number"
    }
    ```
  - **404 Not Found** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 10. **Update User Cash**

**`POST /update-user-cash`**

- **Description:** Updates the cash amount for a user.
- **Request Body:**
  ```json
  {
    "user_id": "number",
    "cash_amount": "number"
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "User cash updated successfully"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

### Leaderboard and Badges

#### 11. **Get Leaderboard**

**`GET /leaderboard`**

- **Description:** Retrieves the leaderboard with user scores.
- **Responses:**
  - **200 OK**
    ```json
    {
      "leaderboard": "leaderboard data"
    }
    ```
  - **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 12. **Get All User Badges**

**`GET /user/:userId/badges`**

- **Description:** Retrieves all badges for a user.
- **Path Parameter:**
  - `userId` - The ID of the user.
- **Responses:**
  - **200 OK**
    ```json
    {
      "badge_name": {
        "description": "string",
        "earned": "boolean"
      }
    }
    ```
  - **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 13. **Update User Badges**

**`POST /update-user-badges`**

- **Description:** Updates badges for a user.
- **Request Body:**
  ```json
  {
    "user_id": "number",
    "badge_ids": ["number"]
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "message": "Badges updated successfully"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

#### 14. **Get Badge Names**

**`POST /get-badge-names`**

- **Description:** Retrieves badge names based on badge IDs.
- **Request Body:**
  ```json
  {
    "badge_ids": ["number"]
  }
  ```
- **Responses:**
  - **200 OK**
    ```json
    {
      "badge_names": "badge names"
    }
    ```
  - **400 Bad Request** or **500 Internal Server Error**
    ```json
    {
      "error": "error message"
    }
    ```

## Running the API

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/quiz-app-backend.git
   ```

2. **Navigate to the project directory:**
   ```bash
   cd quiz-app-backend
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the API server:**
   ```bash
   go run main.go
   ```

5. **Access the API at** `http://localhost:8080`.
