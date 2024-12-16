// Получаем элементы формы
const registerForm = document.getElementById('register-form');
const usernameInput = document.getElementById('username');
const passwordInput = document.getElementById('password');
const emailInput = document.getElementById('email');

// Функция для регистрации нового пользователя
const registerUser = async (username, password, email) => {
  try {
    const response = await fetch('/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password, email }),
    });

    if (response.ok) {
      const data = await response.json();
      alert(data.message); // Показываем сообщение об успехе
      window.location.href = '/login'; // Перенаправляем на страницу входа
    } else {
      const errorData = await response.json();
      alert(errorData.message); // Показываем сообщение об ошибке
    }
  } catch (error) {
    console.error('Ошибка регистрации:', error);
    alert('An error occurred during registration.');
  }
};

// Обработчик отправки формы
registerForm.onsubmit = (event) => {
  event.preventDefault(); // Предотвращаем перезагрузку страницы при отправке формы
  const username = usernameInput.value.trim();
  const password = passwordInput.value.trim();
  const email = emailInput.value.trim();
  
  // Проверяем, что имя и пароль не пустые
  if (username && password) {
    registerUser(username, password, email); // Вызываем функцию для регистрации
  } else {
    alert('Username and password are required.'); // Уведомляем, если данные не указаны
  }
};
