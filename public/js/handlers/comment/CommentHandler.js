import {api as apiUser} from "../../api/users.js";
import {api as apiPost} from "../../api/posts.js";

export default async function FetchComments() {
   try {
       const users = await apiUser.getUsers();
       const userMap = Object.fromEntries(users.map(user => [user.id, user.username]));

       const match = window.location.pathname.match(/\/post\/(\d+)/);
       const postId =  parseInt(match[1]) || null;

       const comments = await apiPost.getComments(postId);

       const container = document.getElementById('comments-container');
       container.innerHTML = ''; // clear previous

       comments.forEach(comment => {
           const authorName = userMap[comment.author_id] || 'Unknown';

           const commentEl = document.createElement('div');
           commentEl.id = `comment-${comment.id}`;
           commentEl.innerHTML = `
           <p><strong>${authorName}</strong>: ${comment.body}</p> ${comment.time}`;
           container.appendChild(commentEl);
       });
   } catch {
       console.log("error");
   }
}