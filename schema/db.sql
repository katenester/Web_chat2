-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS Users (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     username TEXT NOT NULL UNIQUE,
                                     password TEXT NOT NULL,
                                     created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы чатов (индивидуальные чаты между двумя пользователями)
CREATE TABLE IF NOT EXISTS Chats (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     user1_id INTEGER NOT NULL,
                                     user2_id INTEGER NOT NULL,
                                     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                     FOREIGN KEY (user1_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (user2_id) REFERENCES Users(id) ON DELETE CASCADE,
    CONSTRAINT unique_users UNIQUE (user1_id, user2_id)
    );

-- Создание таблицы сообщений
CREATE TABLE IF NOT EXISTS Messages (
                                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        chat_id INTEGER NOT NULL,
                                        sender_id INTEGER NOT NULL,
                                        message TEXT NOT NULL,
                                        sent_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                        FOREIGN KEY (chat_id) REFERENCES Chats(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES Users(id) ON DELETE CASCADE
    );
