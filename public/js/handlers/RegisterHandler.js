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
    }