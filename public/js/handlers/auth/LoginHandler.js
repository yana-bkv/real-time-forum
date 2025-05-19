import {api} from '../../api/users.js';

export default function LoginFormHandler(element) {
    if (!element) return;

    console.log("Login form detected, adding event listener!");

    element.addEventListener("submit", async function (event) {
        event.preventDefault(); // Prevent page reload

        console.log("Login form submitted!"); // Debugging - Check if this appears in Console

        const email = document.getElementById("emailLogin").value;
        const password = document.getElementById("passwordLogin").value;
        var data = {}

        if (email.includes("@")) {
            data = {
                email: email,
                password: password
            };
        } else {
            data = {
                username: email,
                password: password
            };
        }

        try {
            await api.loginUser(data);

            window.location.href = "/login";
        } catch (error) {
            console.error("Error:", error);
        }

    });
    }