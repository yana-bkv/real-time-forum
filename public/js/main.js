document.addEventListener("DOMContentLoaded", function () {
    // Watch for content changes
    const contentDiv = document.getElementById("content");

    // MutationObserver waits for #content to change
    const observer = new MutationObserver(() => {
        const formElement = document.getElementById("register-form");
        if (formElement) {
            console.log("Register form detected, adding event listener!");

            formElement.addEventListener("submit", async function (event) {
                event.preventDefault(); // Prevent page reload

                console.log("Form submitted!"); // Debugging - Check if this appears in Console

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
                } catch (error) {
                    console.error("Error:", error);
                }
            });

            // Stop observing once the form is found
            observer.disconnect();
        }
    });

    observer.observe(contentDiv, { childList: true, subtree: true });
});