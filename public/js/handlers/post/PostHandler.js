import {api} from '../../api/posts.js';
import {api as apiUser } from '../../api/users.js';
import {api as apiTag} from '../../api/tag.js';
import {CreateCommentSection} from '../comment/CreateCommentHandler.js';
import {authService} from "../../helpers/AuthServiceCheck.js";

export default async function postElement(element){
    if (!element) return;

    // Get post ID from URL
    const postId = await window.location.pathname.match(/\/post\/(\d+)/)?.[1];
    if (!postId) {
        console.error("Post ID not found in URL");
        element.textContent = "Invalid post ID";
        return;
    }
    // Get tags by post id
    const tags = await apiTag.getTagsByPostId(postId);

    try {
        const users = await apiUser.getUsers();
        const userMap = Object.fromEntries(users.map(user => [user.id, user.username]));

        const post = await api.getPostById(postId);

        const authorName = userMap[post.author_id] || 'Unknown';

        let categoryHTML = "";
        tags.forEach(tag => {
            categoryHTML += `<p class="post-categories">${tag}</p>`;
        });

        element.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.body}</p>
            <p class="post-author">${authorName}</p>
             ${categoryHTML}
            <p class="post-time">${post.time}</p>
        `;

        // DELETION PROCESS
        const deleteBtn = document.createElement('button');
        deleteBtn.textContent = 'Delete';

        deleteBtn.addEventListener("click", () => {
            console.log("Delete Post");
            api.deletePostById(postId);
            window.location.href = '/posts'
        })

        if (authService.getUsername() === authorName) {
            element.appendChild(deleteBtn);
        }

    } catch (error) {
        console.log(error);
        element.textContent = "Error fetching post";
    }

    const commentSection = document.getElementById("comments-container");
    const commentComponent = CreateCommentSection(postId);
    commentSection.appendChild(commentComponent);
}

