export default function CreatePostHandler(element) {
    if (!element) return;

    console.log("Post form detected, adding event listener!");

    element.addEventListener("submit", async function (event) {
        event.preventDefault(); // Prevent page reload

        console.log("Post form submitted!"); // Debugging - Check if this appears in Console

        const title = document.getElementById("postTitle").value;
        const postBody = document.getElementById("postBody").value;

        const data = {
            title: title,
            body: postBody
        };

        try {
            const response = await fetch("http://localhost:8080/api/post", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data),
                credentials: "include"
            });

            const result = await response.json();
            console.log("Response from server:", result);
            if (response.ok) {
                console.log("Post successfully created!");
                window.location.href = "/posts";
            } else {
                console.log("Post creation failed.");
            }
        } catch (error) {
            console.error("Error:", error);
        }
    });
}