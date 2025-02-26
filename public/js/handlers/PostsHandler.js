export default function postsFeed(element) {
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