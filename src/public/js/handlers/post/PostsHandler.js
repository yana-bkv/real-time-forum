import {api} from '../../api/posts.js';
import {api as apiUser} from "../../api/users.js";
import {LikeHandler} from "../like/LikeHandler.js";

export default async function postsFeed(element){
    if (!element) return;

    try {
        const users = await apiUser.getUsers();
        const userMap = Object.fromEntries(users.map(user => [user.id, user.username]));

        const result = await api.getPosts();

        element.innerHTML = "";

        result.forEach(post => {
            const authorName = userMap[post.author_id] || 'Unknown';

            const postElement = document.createElement("div");
            postElement.classList.add("post-item");
            postElement.innerHTML = `
                    <h3 class="post-title">${post.title}</h3>
                    <p>${post.body}</p>
                    <p>Author: ${authorName}</p>
                    <p>Categories ${post.category}</p>
                    <p>${post.time}</p> 
                `;
            postElement.querySelector(".post-title").addEventListener("click", (event) => {
                event.stopPropagation(); // Prevent bubbling up if needed
                window.location.href = "/post/" + post.id;
            });

            const like = async function () {
                const likeBtn = await LikeHandler(post.id);
                postElement.appendChild(likeBtn);
                element.appendChild(postElement);
            }
            like();
        });

    } catch (error) {
        console.log(error);
        element.textContent = "Error fetching user";
    }
}