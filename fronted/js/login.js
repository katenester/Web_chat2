// Получаем элементы формы
const loginForm = document.getElementById('login-form');
const usernameInput = document.getElementById('username');
const passwordInput = document.getElementById('password');

// Функция для аутентификации пользователя
const loginUser = async (username, password) => {
  try {
    const response = await fetch('/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (response.ok) {
      const data = await response.json();
      // Сохраняем токен в cookie или localStorage для будущих запросов
      document.cookie = `token=${data.token}; path=/`;
      // Перенаправляем на главную страницу
      window.location.href = '/chats'; // Укажите правильный путь к главной странице
    } else {
      const errorData = await response.json();
      alert(errorData.error); // Показываем сообщение об ошибке
    }
  } catch (error) {
    console.error('Ошибка входа:', error);
    alert('An error occurred during login.');
  }
};

// Обработчик отправки формы
loginForm.onsubmit = (event) => {
  event.preventDefault(); // Предотвращаем перезагрузку страницы при отправке формы
  const username = usernameInput.value.trim();
  const password = passwordInput.value.trim();
  
  // Проверяем, что поля не пустые
  if (username && password) {
    loginUser(username, password); // Вызываем функцию входа
  } else {
    alert('Username and password are required.'); // Напоминаем пользователю ввести данные
  }
};
