import {api} from "../../api/users.js";
import {authService} from "../../helpers/AuthServiceCheck.js";

export default async function showUsers(element) {
    if (!element) return;

    const authUser = authService.getUsername();

    console.log("Users list page loaded");

    try {
        const result = await api.getUsers();

        element.innerHTML = "";

        result.forEach(user => {
            const userElement = document.createElement("div");
            userElement.classList.add("user-item");
            userElement.innerHTML = `
                    <h3 class="user-name">Username: ${user.username}</h3>
                    <p>User email: ${user.email}</p>
                    <p>User id: ${user.id}</p>
                    <p>Online status: </p>
                    <a href="/ws/${authUser}/${user.username}">Send a message</a>
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