import {NavigateTo, bindLinkEvents} from '../index.js';

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
            const response = await fetch("http://localhost:8080/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data),
                credentials: "include"
            });

            const result = await response.json();
            console.log("Response from server:", result);
            if (result === "Success") {
                console.log("Login successfully!");
                NavigateTo("/");
                // window.location.href = "/login";
            } else {
                console.log("Login failed.");
            }
        } catch (error) {
            console.error("Error:", error);
        }

        bindLinkEvents();
    });
    }