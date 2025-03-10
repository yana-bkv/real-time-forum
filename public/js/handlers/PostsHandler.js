export default async function postsFeed(element){
    if (!element) return;

    try {
        const response = await fetch("http://localhost:8080/api/posts", {credentials: "include"});

        if (!response.ok) {
            if (response.status === 401) {
                console.log("Unauthorized user");
                return;
            }
            console.log("Response status" + response.status);
            return
        }

        const result = await response.json();
        element.innerHTML = "";

        result.forEach(post => {
            const postElement = document.createElement("div");
            postElement.classList.add("post-item");
            postElement.innerHTML = `
                    <h3 class="post-title">${post.title}</h3>
                    <p>${post.body}</p>
                    <p>${post.author_id}</p>
                    <p>Categories ${post.category}</p>
                    <p>${post.time}</p> 
                `;
            postElement.querySelector(".post-title").addEventListener("click", (event) => {
                event.stopPropagation(); // Prevent bubbling up if needed
                window.location.href = "/post/" + post.id;
            });

            element.appendChild(postElement);
        });
    } catch (error) {
        console.log(error);
        element.textContent = "Error fetching user";
    }
}