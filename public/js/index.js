import PostsForm from './views/Posts.js';
import LoginForm from './views/Login.js';
import RegisterForm from './views/Register.js';
import UserAbout from './views/User.js';

import LoginHandler from './handlers/LoginHandler.js';
import RegisterHandler from './handlers/RegisterHandler.js';
// import PostsHandler from './handlers/PostsHandler.js';
import UserHandler from './handlers/UserHandler.js';
import LogoutHandler from './handlers/LogoutHandler.js';

const navigateTo = url => {
    history.pushState(null, null, url);
    router();
}

const router = async () => {
    const routes = [
        { path: "/", view: PostsForm },
        { path: "/login", view: LoginForm },
        { path: "/register", view: RegisterForm },
        { path: "/user", view: UserAbout },
    ];

    const potentialRoutes = routes.map(route => {
        return {
            route: route,
            isMatch: location.pathname === route.path,
        }
    })

    let match = potentialRoutes.find(potentialMatch => potentialMatch.isMatch);

    if (!match) {
        match = {
            route: routes[0],
            isMatch: true,
        }
    }

    const view = new match.route.view();

    document.querySelector("#content").innerHTML = await view.getHtml();
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    // Routing
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault();
            navigateTo(e.target.href);
        }
    });

    router();

    // Connecting server to client
    const contentDiv = document.getElementById("content");

    const observer = new MutationObserver(() => {
        const loginForm = document.getElementById("loginForm");
        const registerForm = document.getElementById("registerForm");
        // const postForm = document.getElementById("feed");
        const userAbout = document.getElementById("userAbout");
        const logoutButton = document.getElementById("logoutButton");

        LoginHandler(loginForm);
        RegisterHandler(registerForm);
        UserHandler(userAbout);
        LogoutHandler(logoutButton);

        checkAuth();
    })
    observer.observe(contentDiv, { childList: true, subtree: true });

    // Checking if user is registered
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
});



