CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
    cash DECIMAL(10, 2) NOT NULL DEFAULT 100000.00
);

CREATE TABLE quizzes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT
);

CREATE TABLE questions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    quiz_id INT,
    question TEXT NOT NULL,
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);

CREATE TABLE answers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    question_id INT,
    answer TEXT NOT NULL,
    is_correct BOOLEAN,
    FOREIGN KEY (question_id) REFERENCES questions(id)
);

CREATE TABLE user_scores (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    quiz_id INT,
    score INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);

CREATE TABLE user_quizzes (
    user_id INT,
    quiz_id INT,
    completed_at DATETIME,
    PRIMARY KEY (user_id, quiz_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);


CREATE TABLE badges (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,      
    description TEXT,                 
);

CREATE TABLE user_badges (
    user_id INT,
    badge_id INT,
    earned BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (user_id, badge_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (badge_id) REFERENCES badges(id)
);



