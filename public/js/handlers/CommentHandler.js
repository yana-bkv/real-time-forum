const currentUser = {
    id: 1};

export default function FetchComments() {
    let userMap = {};

    fetch('http://localhost:8080/api/users') // backend should return [{id: 1, name: "Alice"}, ...]
        .then(res => res.json())
        .then(users => {
            userMap = Object.fromEntries(users.map(user => [user.id, user.username]));
        });


    const match = window.location.pathname.match(/\/post\/(\d+)/);
    const postId =  parseInt(match[1]) || null;

    fetch(`http://localhost:8080/api/post/${postId}/comments`)
        .then(response => response.json())
        .then(comments => {

            const container = document.getElementById('comments-container');
            container.innerHTML = ''; // clear previous

            comments.forEach(comment => {
                const authorName = userMap[comment.author_id] || 'Unknown';

                const canDelete = currentUser.id === comment.authorId ;

                const commentEl = document.createElement('div');
                commentEl.id = `comment-${comment.id}`;
                commentEl.innerHTML = `
          <p><strong>${authorName}</strong>: ${comment.body}</p> <span style="font-size:12px;">${comment.time}</span>
          ${canDelete ? `<button onclick="deleteComment(${comment.postId}, ${comment.id})">Delete</button>` : ''}
        `;
                container.appendChild(commentEl);
            });
        })
        .catch(err => {
            console.error('Failed to load comments:', err);
        });
}