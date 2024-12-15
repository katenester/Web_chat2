# Web_chat2

1. Регистрация нового пользователя. POST /auth/register
Запрос: json
{
"username": "user1",
"password": "password123"
}
Ответы:
# 201 Created — пользователь успешно зарегистрирован.
json
{
"message": "User registered successfully"
}
# 400 Bad Request — имя пользователя уже существует или данные неверны.
json
{
"message": "Username already exists or invalid data"
}
# 500 StatusInternalServerError — ошибка на сервере
{
"message": "User registered successfully"
}
2. POST /auth/login Аутентификация пользователя и получение токена.

Запрос: json
{
"username": "user1",
"password": "password123"
}
Ответы:
# 200 OK — успешная аутентификация, возвращается токен.
json
{
"token": "your-auth-token"
}
# 401 Unauthorized — неправильный логин или пароль.
json
{
"error": "Invalid username or password"
}
# 500 StatusInternalServerError — ошибка на сервере
{
"error": "Invalid username or password"
}
3. Получение списка чатов пользователя. GET /chat
  Заголовок 
  - Bearer Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgxOTUwOTcsImlhdCI6MTcyODE1MTg5NywidXNlcl9pZCI6MX0.ZPqAiu0edno2Z08VEpjNfMmJ2Dg22KhS2ZPSEHaqKDo
Ответы:
# 200 OK — возвращает список чатов.
json
[
{
"chat_id": 1,
"user": "user_name"
}
{
"chat_id": 2,
"user": "user_name2"
}
]
# 401 Unauthorized — пользователь не аутентифицирован.
4. Создание нового чата между двумя пользователями POST /chat/:user_name
   Заголовок
- Bearer Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgxOTUwOTcsImlhdCI6MTcyODE1MTg5NywidXNlcl9pZCI6MX0.ZPqAiu0edno2Z08VEpjNfMmJ2Dg22KhS2ZPSEHaqKDo

Ответы:
# 201 Created — чат успешно создан.
# 400 Bad Request — если чат уже существует или user2_id не существует.
json
{
"error": "Chat already exists or invalid user"
}

5. Отправка сообщения в чат. POST /chat/messages/:user_name
   Заголовок
- Bearer Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgxOTUwOTcsImlhdCI6MTcyODE1MTg5NywidXNlcl9pZCI6MX0.ZPqAiu0edno2Z08VEpjNfMmJ2Dg22KhS2ZPSEHaqKDo
Запрос:
json
{
"message": "Hello, how are you?"
}
Ответы:
# 201 Created — сообщение успешно отправлено.

# 400 Bad Request — некорректный запрос, например, если сообщение пустое или чат не существует.
json
{
"error": "Invalid message or chat not found"
}
# 401 Unauthorized — пользователь не аутентифицирован.

6. Получение всех сообщений из чата. GET /chat/messages/:user_name
   Заголовок
- Bearer Token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjgxOTUwOTcsImlhdCI6MTcyODE1MTg5NywidXNlcl9pZCI6MX0.ZPqAiu0edno2Z08VEpjNfMmJ2Dg22KhS2ZPSEHaqKDo
Ответы:
# 200 OK — возвращает список сообщений из чата.
json
[
{
"sender": "user_name"",
"message": "Hello, how are you?"
}
{
"sender": "user_name"",
"message": "Good"
}
]
# 404 Not Found — чат не найден.
json
{
"error": "Chat not found"
}
# 401 Unauthorized — пользователь не аутентифицирован.

