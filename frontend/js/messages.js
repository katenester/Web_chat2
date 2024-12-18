document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    const username =localStorage.getItem('friend');
    // const username = window.location.pathname.split('/').pop();  // Get username from URL
    const chatUsernameElement = document.getElementById('chatUsername');
    const chatBox = document.getElementById('chatBox');
    const messageInput = document.getElementById('messageInput');

    if (!token) {
        alert('Unauthorized access');
        window.location.href = '/login';
        return;
    }

    chatUsernameElement.textContent = username;

    // Fetch messages when page loads
    fetch(`/api/chat/messages/${username}`, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
        .then(response => {
            if (response.status === 401) {
                alert('Unauthorized');
                window.location.href = '/login';
            } else if (response.status === 404) {
                alert('Chat not found');
            } else {
                return response.json();
            }
        })
        .then(data => {
            if (Array.isArray(data)) {
                data.forEach(msg => {
                    const msgDiv = document.createElement('div');
                    msgDiv.classList.add('message');
                    const isSender = msg.sender === username;

                    msgDiv.classList.add(isSender ? 'message-left' : 'message-right');
                    msgDiv.innerHTML = `
                        <span class="sender">${msg.sender}</span>
                        <span class="content">${msg.message}</span>
                    `;
                    chatBox.appendChild(msgDiv);
                });
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred. Please try again.');
        });

    // Send message function
    window.sendMessage = function() {
        const message = messageInput.value.trim();
        if (!message) {
            alert('Message cannot be empty');
            return;
        }

        fetch(`/api/chat/messages/${username}`, {
            method: 'POST',
            headers: {
                'Authorization': 'Bearer ' + token,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ message: message })
        })
            .then(response => response.json())
            .then(data => {
                if (data.status === 'success') {
                    const msgDiv = document.createElement('div');
                    msgDiv.classList.add('message', 'message-right');
                    msgDiv.innerHTML = `
                        <span class="sender">You</span>
                        <span class="content">${message}</span>
                    `;
                    chatBox.appendChild(msgDiv);
                    messageInput.value = '';  // Clear the input field
                    chatBox.scrollTop = chatBox.scrollHeight;  // Scroll to the bottom
                } else {
                    alert(data.error || 'Failed to send message');
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('An error occurred while sending the message');
            });
    };
});