import {api} from '../../api/posts.js';

export default function CreatePostHandler(element) {
    if (!element) return;

    console.log("Post form detected, adding event listener!");

    element.addEventListener("submit", async function (event) {
        event.preventDefault(); // Prevent page reload

        console.log("Post form submitted!"); // Debugging - Check if this appears in Console

        const title = document.getElementById("postTitle").value;
        const postBody = document.getElementById("postBody").value;

        const data = {
            title: title,
            body: postBody
        };

        try {
            const result = await api.createPost(data);
            console.log("Post successfully created", result);
            window.location.href = "/posts";
        } catch (error) {
            console.error("Error:", error);
        }
    });
}