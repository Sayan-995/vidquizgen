const URL="http://localhost:8080"
document.addEventListener('DOMContentLoaded', function() {
    // Elements
    const videoUrlInput = document.getElementById('videoUrl');
    const searchBtn = document.getElementById('searchBtn');
    const loader = document.getElementById('loader');
    const results = document.getElementById('results');
    const problemsList = document.getElementById('problemsList');
    const error = document.getElementById('error');
    const errorMessage = document.getElementById('errorMessage');
    const exampleLink = document.querySelector('.example-link');
    
    // Event listeners
    searchBtn.addEventListener('click', searchProblems);
    videoUrlInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            searchProblems();
        }
    });
    
    // Set example video URL when clicked
    exampleLink.addEventListener('click', function(e) {
        e.preventDefault();
        videoUrlInput.value = this.textContent;
    });
    
    // Search for problems
    function searchProblems() {
        const videoUrl = videoUrlInput.value.trim();
        
        if (!videoUrl) {
            showError('Please enter a YouTube URL');
            return;
        }
        
        if (!isValidYouTubeUrl(videoUrl)) {
            showError('Please enter a valid YouTube URL');
            return;
        }
        
        // Show loader and hide other sections
        loader.classList.remove('hidden');
        results.classList.add('hidden');
        error.classList.add('hidden');
        
        // Send request to API
        fetch(`${URL}/api/quiz`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url: videoUrl })
        })
        .then(response => {
            if (!response.ok) {
                return response.json().then(data => {
                    throw new Error(data.error || 'Failed to get problems');
                });
            }
            return response.json();
        })
        .then(data => {
            displayResults(data.questions);
        })
        .catch(err => {
            showError(err.message);
        })
        .finally(() => {
            loader.classList.add('hidden');
        });
    }
    
    // Display results
    function displayResults(questions) {
        // Clear previous results
        problemsList.innerHTML = '';
        
        if (!questions || questions.length === 0) {
            showError('No relevant problems found for this video');
            return;
        }
        
        // Create and append problem cards
        questions.forEach((question, index) => {
            const card = document.createElement('div');
            card.className = 'problem-card fade-in';
            card.style.animationDelay = `${index * 0.1}s`;
            
            card.innerHTML = `
                <h3 class="problem-title">${question.title}</h3>
                <a href="${question.url}" class="problem-link" target="_blank">Solve on LeetCode</a>
            `;
            
            problemsList.appendChild(card);
        });
        
        // Show results section
        results.classList.remove('hidden');
    }
    
    // Show error message
    function showError(message) {
        errorMessage.textContent = message;
        error.classList.remove('hidden');
        results.classList.add('hidden');
        loader.classList.add('hidden');
    }
    
    // Validate YouTube URL
    function isValidYouTubeUrl(url) {
        const regex = /^(https?:\/\/)?(www\.)?(youtube\.com|youtu\.?be)\/.+/;
        return regex.test(url);
    }
    
    // Add some fun interactivity
    const emoji = ['🚀', '💻', '🔥', '⚡️', '🧠', '💡'];
    
    function randomEmoji() {
        return emoji[Math.floor(Math.random() * emoji.length)];
    }
    
    // Easter egg: Konami code
    let konamiIndex = 0;
    const konamiCode = [38, 38, 40, 40, 37, 39, 37, 39, 66, 65]; // Up, Up, Down, Down, Left, Right, Left, Right, B, A
    
    document.addEventListener('keydown', function(e) {
        if (e.keyCode === konamiCode[konamiIndex]) {
            konamiIndex++;
            
            if (konamiIndex === konamiCode.length) {
                activatePartyMode();
                konamiIndex = 0;
            }
        } else {
            konamiIndex = 0;
        }
    });
    
    function activatePartyMode() {
        document.body.classList.add('party-mode');
        
        // Add floating emojis
        for (let i = 0; i < 20; i++) {
            const emojiSpan = document.createElement('span');
            emojiSpan.className = 'floating-emoji';
            emojiSpan.textContent = randomEmoji();
            emojiSpan.style.left = `${Math.random() * 100}%`;
            emojiSpan.style.animationDuration = `${3 + Math.random() * 7}s`;
            emojiSpan.style.animationDelay = `${Math.random() * 2}s`;
            document.body.appendChild(emojiSpan);
            
            // Remove after animation
            setTimeout(() => {
                emojiSpan.remove();
            }, 10000);
        }
        
        // Remove party mode after 10 seconds
        setTimeout(() => {
            document.body.classList.remove('party-mode');
        }, 10000);
    }
    
    // Add party mode CSS
    const partyModeStyle = document.createElement('style');
    partyModeStyle.textContent = `
        .party-mode {
            overflow: hidden;
        }
        
        .floating-emoji {
            position: fixed;
            top: -10%;
            font-size: 2rem;
            animation: float linear forwards;
            z-index: 1000;
        }
        
        @keyframes float {
            0% {
                top: -10%;
                transform: translateX(0) rotate(0deg);
            }
            100% {
                top: 110%;
                transform: translateX(100px) rotate(360deg);
            }
        }
    `;
    document.head.appendChild(partyModeStyle);
});