export default async function showUsers(element) {
    if (!element) return;

    console.log("Users list page loaded");

    try {
        const response = await fetch("http://localhost:8080/api/users", {
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

        element.innerHTML = "";

        result.forEach(user => {
            const userElement = document.createElement("div");
            userElement.classList.add("user-item");
            userElement.innerHTML = `
                    <h3 class="user-name">Username: ${user.username}</h3>
                    <p>User email: ${user.email}</p>
                    <p>User id: ${user.id}</p>
                    <p>Online status: </p>
                `;
            userElement.querySelector(".user-name").addEventListener("click", (event) => {
                event.stopPropagation(); // Prevent bubbling up if needed
                window.location.href = "/user/" + user.id;
            });

            element.appendChild(userElement);
        });

    } catch (error) {
        console.error("Error fetching user:", error);
        element.textContent = "Error fetching user";
    }
}