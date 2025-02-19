document.addEventListener("DOMContentLoaded", function () {
    // Watch for content changes
    const contentDiv = document.getElementById("content");

    // MutationObserver waits for #content to change
    const observer = new MutationObserver(() => {
        const registerForm = document.getElementById("register-form");
        const loginForm = document.getElementById("login-form");
        const userHello = document.getElementById("userHello");

        if (registerForm) {
            registerFormHandler(registerForm);
        }

        if (loginForm) {
            loginFormHandler(loginForm);
        }

        if (userHello) {
            showUser(userHello);
        }

        // Stop observing only when both forms are found
        if (registerForm && loginForm) {
            observer.disconnect();
        }
    });

    observer.observe(contentDiv, { childList: true, subtree: true });
});

// Register form
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
                    console.log("Register successfully!");
                    // redirect to login page after successfully registration
                    window.location.href = "/login";
                } else {
                    console.log("Register failed.");
                }
            } catch (error) {
                console.error("Error:", error);
            }
        });
    }}

// Login form
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
                    body: JSON.stringify(data),
                    credentials: "include"
                });

                const result = await response.json();
                console.log("Response from server:", result);
                if (result === "Success") {
                    // redirect to login page after successfully registration
                    console.log("Login successfully!");
                    window.location.href = "/";
                } else {
                    console.log("Login failed.");
                }
            } catch (error) {
                console.error("Error:", error);
            }
        });
    }}


function showUser(element) {
    (async() => {
        try {
            const response = await fetch("http://localhost:8080/api/user", {
                headers: {
                    "Content-Type": "application/json"
                },
                credentials: "include"
            });

            if (!response.ok) {
                if (response.status === 401) {
                    console.log("Unauthorized user");
                    element.textContent = "You are not logged in";
                    return;
                }
                throw new Error(response.status);
            }

            const result = await response.json();
            element.textContent = result.username ? `Hi ${result.username}` : "You are not logged in";

        } catch (error) {
            console.log(error);
            element.textContent = "Error fetching user";
        }
   })();

}