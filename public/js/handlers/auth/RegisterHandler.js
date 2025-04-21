import {api} from "../../api/users.js";

export default function registerFormHandler(element) {
    if (!element) return;

    console.log("Register form detected, adding event listener!");

    element.addEventListener("submit", async function (event) {
        event.preventDefault(); // Prevent page reload

        console.log("Register form submitted!"); // Debugging - Check if this appears in Console

        const username = document.getElementById("username").value;
        const email = document.getElementById("email").value;
        const password = document.getElementById("password").value;

        const data = {
            username: username,
            email: email,
            password: password
        };

        try {
            await api.registerUser(data);
            window.location.href = "/login";
        } catch (error) {
            console.error("Error:", error);
        }
    });
    }