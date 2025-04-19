export default async function postElement(element){
    if (!element) return;

    // Get post ID from URL
    const postId = await window.location.pathname.match(/\/post\/(\d+)/)?.[1];

    if (!postId) {
        console.error("Post ID not found in URL");
        element.textContent = "Invalid post ID";
        return;
    }

    try {
        let userMap = {};

        fetch('http://localhost:8080/api/users') // backend should return [{id: 1, name: "Alice"}, ...]
            .then(res => res.json())
            .then(users => {
                userMap = Object.fromEntries(users.map(user => [user.id, user.username]));
            });

        const response = await fetch(`http://localhost:8080/api/post/${postId}`, {credentials: "include"});

        if (!response.ok) {
            if (response.status === 401) {
                console.log("Post not found with status code");
                return;
            }
            console.log("Response status" + response.status);
            return
        }

        const post = await response.json();
        const authorName = userMap[post.author_id] || 'Unknown';
        element.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.body}</p>
            <p class="post-author">${authorName}</p>
            <p class="post-categories">Categories ${post.category}</p>    
            <p class="post-time">${post.time}</p> 
        `;

        // DELETION PROCESS
        const deletePostBtn = document.getElementById("deletePost");

        deletePostBtn.addEventListener("click", () => {
            console.log("Delete Post");
            console.log(post.author_id);
            //deletePost(postId);
        })

    } catch (error) {
        console.log(error);
        element.textContent = "Error fetching post";
    }
}

function deletePost(postId){
    fetch(`http://localhost:8080/api/post/${postId}`,
        {credentials: "include",
            method: 'DELETE',
        })
        .then(response => {
            if (response.ok) {
                alert(`Post deleted successfully`);
                window.location.href = '/posts';
            } else {
                alert('failed to delete post');
            }
        })
        .catch(error => {
            console.log(error);
        })
}