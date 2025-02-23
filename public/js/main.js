document.addEventListener("DOMContentLoaded", function () {
    // Watch for content changes
    const contentDiv = document.getElementById("content");

    // MutationObserver waits for #content to change
    const observer = new MutationObserver(() => {
        // const path = window.location.pathname.split('/').filter(Boolean).pop();

        const registerForm = document.getElementById("registerForm");
        const loginForm = document.getElementById("loginForm");
        const hiUserP = document.getElementById("hiUser");
        const feed = document.getElementById("feed");

        if (registerForm) {
            registerFormHandler(registerForm);
        }

        if (loginForm) {
            loginFormHandler(loginForm);
        }

        if (hiUserP) {
            showUser(hiUserP);
        }

        if (feed) {
            postsFeed(feed);
        }

        // Stop observing only when both forms are found
        if (registerForm && loginForm && hiUserP && feed) {
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
                console.log("Response status" + response.status);
            }

            const result = await response.json();

            element.textContent = result.username ? `Hi ${result.username}` : "You are not logged in";

        } catch (error) {
            console.log(error);
            element.textContent = "Error fetching user";
        }
   })();

}

async function logout() {
    console.log("Pressed logout button");
    const response = await fetch("http://localhost:8080/api/logout", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        credentials: "include"
    });

    if (!response.ok) {
        console.log("Response status" + response.status);
    }

    const result = await response.json();
    console.log(result);
    location.reload();
}

function postsFeed(element) {
    (async() => {
        try {
            const response = await fetch("http://localhost:8080/api/user", {credentials: "include"});

            if (!response.ok) {
                if (response.status === 401) {
                    console.log("Unauthorized user");
                    return;
                }
                console.log("Response status" + response.status);
            }

            const result = await response.json();
            element.textContent = result.username ? `Posts for ${result.username}` : "You are not logged in";

        } catch (error) {
            console.log(error);
            element.textContent = "Error fetching user";
        }
    })();

}
