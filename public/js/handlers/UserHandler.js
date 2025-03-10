export default async function showUser(element) {
    if (!element) return;

    console.log("User page loaded");

        try {
            const response = await fetch("http://localhost:8080/api/user", {
                headers: {
                    "Content-Type": "application/json"
                },
                credentials: "include"
            });

            console.log("Fetching user data...");

            if (!response.ok) {
                if (response.status === 401) {
                    console.log("Unauthorized user");
                    return;
                }
                console.log("Response status" + response.status);
                element.textContent = "Error fetching user";
                return;
            }

            const result = await response.json();

            const userNameEl = document.querySelector("p.userName");
            const userEmailEl = document.querySelector("p.userEmail");

            if (result.username) {
                if (userNameEl) userNameEl.textContent = "Username: " + result.username;
                if (userEmailEl) userEmailEl.textContent = "Email: " +  result.email;
            } else {
                if (userNameEl) userNameEl.textContent = "You are not logged in";
            }

        } catch (error) {
            console.error("Error fetching user:", error);
            element.textContent = "Error fetching user";
        }
}