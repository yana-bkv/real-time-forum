import {api as apiUser} from "../../api/users.js";
import {api as apiMsg} from "../../api/comments.js";
import {authService} from "../../helpers/AuthServiceCheck.js";

export default async function FetchComments() {
   try {
       const users = await apiUser.getUsers();
       const userMap = Object.fromEntries(users.map(user => [user.id, user.username]));

       const match = window.location.pathname.match(/\/post\/(\d+)/);
       const postId =  parseInt(match[1]) || null;

       const comments = await apiMsg.getComments(postId);

       const container = document.getElementById('comments-container');
       container.innerHTML = ''; // clear previous

       comments.forEach(comment => {
           const authorName = userMap[comment.author_id] || 'Unknown';

           const commentEl = document.createElement('div');
           const deleteBtn = document.createElement('button');
           deleteBtn.textContent = "Delete";
           commentEl.id = `comment-${comment.id}`;
           commentEl.innerHTML = `
            <p class="msg-receiver"><span class="msg-title">${authorName}</span>: ${comment.body} <span class="msg-timestamp">(${comment.time})</span> </p>
            `;
           deleteBtn.addEventListener('click', () => {
               deleteComment(postId, comment.id);
           })
           if (authService.getUsername() === authorName) {
               commentEl.querySelector('.msg-receiver').appendChild(deleteBtn);
           }
           container.appendChild(commentEl);
       });
   } catch {
       console.log("error");
   }
}

async function deleteComment(postId, commentId) {
    await apiMsg.deleteComment(postId, commentId);
}