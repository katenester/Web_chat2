alert('Привет');

// Получаем элементы из DOM
const loginButton = document.getElementById('login-button');
const signupButton = document.getElementById('signup-button');
const userBox = document.getElementById('user-box');
const usernameDisplay = document.getElementById('username-display');
const signoutLink = document.getElementById('signout-link');
const chatSetup = document.getElementById('chat-setup');

// Функция для проверки авторизации
const checkAuth = () => {
    const token = getCookie('token'); // Получаем токен из cookies

    if (token) {
        // Если токен есть, показываем информацию о пользователе
        const username = getCookie('username'); // Получаем имя пользователя из cookies
        usernameDisplay.textContent = username;
        userBox.style.display = 'flex'; // Показываем информацию о пользователе
        chatSetup.style.display = 'block'; // Показываем секцию чата
        loginButton.style.display = 'none'; // Скрываем кнопки входа
        signupButton.style.display = 'none'; // Скрываем кнопки регистрации
    } else {
        // Если токена нет, показываем кнопки логина и регистрации
        userBox.style.display = 'none';
        chatSetup.style.display = 'none';
        loginButton.style.display = 'block';
        signupButton.style.display = 'block';
    }
};

// Вход пользователя
loginButton.onclick = () => {
    window.location.href = '/login'; // Перенаправляем на страницу входа
};

// Регистрация пользователя
signupButton.onclick = () => {
    window.location.href = '/register'; // Перенаправляем на страницу регистрации
};

// Выход пользователя
signoutLink.onclick = (e) => {
    e.preventDefault(); // Предотвращаем переход по ссылке
    document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'; // Удаляем токен
    document.cookie = 'username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'; // Удаляем имя пользователя тоже
    checkAuth(); // Обновляем интерфейс
};

// Функция для получения cookie по имени
function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

// Проверяем авторизацию при загрузке страницы
checkAuth();
