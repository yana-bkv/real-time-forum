export default function LoginFormHandler(element) {
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