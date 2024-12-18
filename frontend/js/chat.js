document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    if (!token) {
        alert('Unauthorized access');
        window.location.href = '/login.html';
        return;
    }

    // Fetch the list of chats
    fetch('/api/chat', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
        .then(response => {
            if (response.status === 401) {
                alert('Unauthorized');
                window.location.href = '/login';
                return;
            } else if (response.status !== 200) {
                alert('Failed to fetch chats. Server responded with status ' + response.status);
                return Promise.reject('Failed to fetch chats');
            } else {
                return response.json();
            }
        })
        .then(data => {
            //console.log('Received data:', data); // Отладка: выводим полученные данные
            const chatList = document.getElementById('chatList');
            chatList.innerHTML = ''; // Очищаем список перед добавлением новых чатов

            if (Array.isArray(data)) {
                data.forEach(chat => {
                    const li = document.createElement('li');
                    li.textContent = `Chat with ${chat.user}`;
                    li.onclick = function() {
                        localStorage.setItem('friend', chat.user);
                        window.location.href = `/messages`;
                    };
                    chatList.appendChild(li);
                });
            } else {
                alert('Unexpected response format');
                console.error('Expected array, received:', data);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred. Please try again.'+error);
        });
});

// Toggle the visibility of the "Create New Chat" form
function toggleCreateChatForm() {
    const form = document.getElementById('createChatForm');
    form.style.display = form.style.display === 'block' ? 'none' : 'block';
}

// Create a new chat
function createChat() {
    const token = localStorage.getItem('token');
    const username = document.getElementById('newChatUsername').value.trim();

    if (!username) {
        alert('Please enter a valid username');
        return;
    }

    fetch(`/api/chat/${username}`, {
        method: 'POST',
        headers: {
            'Authorization': 'Bearer ' + token,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ user_name: username })
    })
        .then(response => {
            alert('Chat created successfully');
            localStorage.setItem('friend', username);
            window.location.href = `/messages`;

        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while creating the chat');
        });
}