CREATE TABLE Todos (
    todo_id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(user_id), 
    title VARCHAR(1024) NOT NULL,
    description TEXT,
    creation_time TIME,
    due_time TIMESTAMP NULL,
    complete_time TIMESTAMP NULL,
    completion_comment TEXT
);