import {api} from '../../api/users.js';

export default async function showUser(element) {
    if (!element) return;

    console.log("User page loaded");

        try {
            const result = await api.getAuthUser();

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