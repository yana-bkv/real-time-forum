export default function LogoutHandler() {
    document.getElementById("logoutButton").addEventListener("click", async (event) => {
        event.preventDefault();

        console.log("Log out found");
        try {
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

            window.location.href = "/";
        } catch (error) {
            console.log("Logout failed", error);
        }
    })
}