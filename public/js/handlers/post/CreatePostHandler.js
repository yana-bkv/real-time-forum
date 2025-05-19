import {api} from '../../api/posts.js';
import {CategoryHandler, showCategoryForm, createCategory} from "../tag/CategoryHandler.js";

let isCreatePostHandlerInitialized = false;

export default async function CreatePostHandler(element) {
    if (!element || isCreatePostHandlerInitialized) return;
    isCreatePostHandlerInitialized = true;

    console.log("Post form detected, adding event listener!");

    const showCategoryBtn = document.getElementById("show-category-button");
    const saveCategoryBtn = document.getElementById("save-category-button");

    if (showCategoryBtn && saveCategoryBtn) {
        showCategoryBtn.addEventListener("click", showCategoryForm);
        saveCategoryBtn.addEventListener("click", createCategory);
    }

    console.log("Вызываем CategoryHandler...");
    try {
        await CategoryHandler();
    } catch (error) {
        console.log(error);
    }

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