document.addEventListener('DOMContentLoaded', () => {
    const chatList = document.getElementById('chat-list');
    const addChatBtn = document.getElementById('add-chat-btn');
    const usernameInput = document.getElementById('username');
    const errorMessage = document.getElementById('error-message');

    // Получаем токен из cookie или localStorage
    const token = document.cookie.match(/token=([^;]+)/)?.[1] || localStorage.getItem('token');
    if (!token) {
        window.location.href = '/login.html';  // Перенаправляем на страницу логина, если нет токена
    }

    // Функция для получения списка чатов
    function fetchChats() {
        fetch('/chats', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
            .then(response => response.text()) // Получаем ответ как текст для отладки
            .then(text => {
                console.log('Server Response:', text);
                try {
                    const data = JSON.parse(text);
                    // Дальше обработка данных
                } catch (error) {
                    console.error('Error parsing JSON:', error);
                }
            })
            .catch(error => {
                console.error('Error:', error);
            });

    }

    // Функция для создания нового чата
    function createChat(username) {
        fetch(`/chats/${username}`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else if (response.status === 400) {
                    return response.json().then(data => {
                        throw new Error(data.error);
                    });
                } else {
                    throw new Error('Failed to create chat');
                }
            })
            .then(() => {
                fetchChats(); // Обновляем список чатов
                usernameInput.value = ''; // Очищаем поле ввода
                errorMessage.textContent = '';
            })
            .catch(error => {
                errorMessage.textContent = 'Error creating chat: ' + error.message;
            });
    }

    // Добавляем обработчик на кнопку создания чата
    addChatBtn.addEventListener('click', () => {
        const username = usernameInput.value.trim();
        if (username) {
            createChat(username);
        } else {
            errorMessage.textContent = 'Please enter a username';
        }
    });

    // Загружаем список чатов при загрузке страницы
    fetchChats();
});
