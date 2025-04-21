import {api} from '../../api/users.js';

export default function LogoutHandler() {
    document.getElementById("logoutButton").addEventListener("click", async (event) => {
        event.preventDefault();

        console.log("Log out found");
        try {
            await api.logoutUser();

            window.location.href = "/";
        } catch (error) {
            console.log("Logout failed", error);
        }
    })
}