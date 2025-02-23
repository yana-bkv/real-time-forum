const routes = {
    "/": {
        template: "/templates/posts.html",
    },
    "/login": {
        template: "/templates/login.html",
    },
    "/register": {
        template: "/templates/register.html",
    },
};

function navigate(event, path) {
    event.preventDefault(); // Prevent page reload
    history.pushState({}, "", path); // Change URL without reloading
    router();
}

const router = async () => {
    const content = document.getElementById("content");
    var path = window.location.pathname;
    if (path.length === 0) {
        path = "/";
    }
    const route = routes[path] || "404";
    const html = await fetch(route.template).then((res) => res.text());
    content.innerHTML = html;
    attachNavListeners()
}

// Function to reattach event listeners on all navigation links
function attachNavListeners() {
    document.querySelectorAll("nav a").forEach(link => {
        link.addEventListener("click", function(event) {
            navigate(event, this.getAttribute("href"));
        });
    });
}


document.addEventListener("DOMContentLoaded", function () {
    checkAuth();
    // Watch for content changes
    const nav = document.querySelector("nav");

    // MutationObserver waits for #content to change
    const observer = new MutationObserver(() => {
        // Check if user in logged in
        checkAuth();
    })

    observer.observe(nav, { childList: true, subtree: true });

    async function checkAuth() {
        try {
            const response = await fetch("http://localhost:8080/api/user", { credentials: "include" });
            if (response.ok) {
                document.getElementById("authorized").style.display = "block";
                document.getElementById("unauthorized").style.display = "none";
            } else {
                document.getElementById("authorized").style.display = "none";
                document.getElementById("unauthorized").style.display = "block";
            }
        } catch (error) {
            console.error("Auth check failed:", error);
        }
    }
})

// Attach on initial load
window.addEventListener("load", () => {
    router();
    attachNavListeners(); // 🔥 Attach event listeners on first load
});

window.addEventListener("popstate", router);
document.addEventListener("DOMContentLoaded", router);
