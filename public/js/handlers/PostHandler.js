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
        const response = await fetch(`http://localhost:8080/api/post/${postId}`, {credentials: "include"});

        if (!response.ok) {
            if (response.status === 401) {
                console.log("Unauthorized user");
                return;
            }
            console.log("Response status" + response.status);
            return
        }

        const post = await response.json();
        element.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.body}</p>
            <p class="post-author">${post.author_id}</p>
            <p class="post-categories">Categories ${post.category}</p>    
            <p class="post-time">${post.time}</p> 
        `;
    } catch (error) {
        console.log(error);
        element.textContent = "Error fetching user";
    }
}