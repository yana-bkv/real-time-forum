document.addEventListener("DOMContentLoaded", function () {
    // Watch for content changes
    const contentDiv = document.getElementById("content");

    // MutationObserver waits for #content to change
    const observer = new MutationObserver(() => {
        const registerForm = document.getElementById("register-form");
        const loginForm = document.getElementById("login-form");

        if (registerForm) {
            registerFormHandler(registerForm);
        }

        if (loginForm) {
            loginFormHandler(loginForm);
        }

        // Stop observing only when both forms are found
        if (registerForm && loginForm) {
            observer.disconnect();
        }
    });

    observer.observe(contentDiv, { childList: true, subtree: true });
});

function registerFormHandler(element) {
    if (element) {
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
                const response = await fetch("http://localhost:8080/api/register", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(data)
                });

                const result = await response.json();
                console.log("Response from server:", result);
                if (result === "Success") {
                    alert("Register successfully!");
                    // redirect to login page after succescfull registration
                    window.location.href = "/login";
                } else {
                    alert("Register failed.");
                }
            } catch (error) {
                console.error("Error:", error);
            }
        });
    }}

function loginFormHandler(element) {
    if (element) {
        console.log("Login form detected, adding event listener!");

        element.addEventListener("submit", async function (event) {
            event.preventDefault(); // Prevent page reload

            console.log("Login form submitted!"); // Debugging - Check if this appears in Console

            const email = document.getElementById("emailLogin").value;
            const password = document.getElementById("passwordLogin").value;

            const data = {
                email: email,
                password: password
            };

            try {
                const response = await fetch("http://localhost:8080/api/login", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(data)
                });

                const result = await response.json();
                console.log("Response from server:", result);
            } catch (error) {
                console.error("Error:", error);
            }
        });
    }}