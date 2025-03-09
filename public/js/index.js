import PostsForm from './views/Posts.js';
import PostForm from './views/Post.js';
import LoginForm from './views/Login.js';
import RegisterForm from './views/Register.js';
import UserAbout from './views/User.js';
import WelcomePage from './views/WelcomePage.js';

import LoginHandler from './handlers/LoginHandler.js';
import RegisterHandler from './handlers/RegisterHandler.js';
import PostsHandler from './handlers/PostsHandler.js';
import PostHandler from './handlers/PostHandler.js';
import UserHandler from './handlers/UserHandler.js';
import LogoutHandler from './handlers/LogoutHandler.js';

const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const getParams = match => {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);

    return Object.fromEntries(keys.map((key,i) => {
        return [key,values[i]];
    }));
};

export const NavigateTo = url => {
    console.log(`Navigating to: ${url}`); // Debugging

    history.pushState(null, null, url);
    router();
}

export const bindLinkEvents = () => {
    document.body.addEventListener("click", e => {
        if (e.target.closest("[data-link]")) {
            e.preventDefault();
            NavigateTo(e.target.href);
        }
    })}

const router = async () => {
    const routes = [
        { path: "/", view: WelcomePage },
        { path: "/posts", view: PostsForm },
        { path: "/post/:id", view: PostForm},
        { path: "/login", view: LoginForm },
        { path: "/register", view: RegisterForm },
        { path: "/user", view: UserAbout },
    ];

    const potentialRoutes = routes.map(route => {
        return {
            route: route,
            result: location.pathname.match(pathToRegex((route.path))),
            // isMatch: location.pathname === route.path,
        }
    })

    let match = potentialRoutes.find(potentialMatch => potentialMatch.result !== null);

    if (!match) {
        setTimeout(() => {
            console.log('No match found!');
            // navigateTo("/posts");
        }, 2000);
        return;
    }

    const view = new match.route.view(getParams(match));

    document.querySelector("#content").innerHTML = await view.getHtml();
};

window.addEventListener("popstate", router);

let postsFetched = false; // Track if posts are already fetched
let postFetched = false; // Track if posts are already fetched

document.addEventListener("DOMContentLoaded", () => {
    // Routing
    bindLinkEvents();
    router();

    // Connecting server to client
    const contentDiv = document.getElementById("content");

    const observer = new MutationObserver(() => {
        const loginForm = document.getElementById("loginForm");
        const registerForm = document.getElementById("registerForm");
        const userAbout = document.getElementById("userAbout");
        const logoutButton = document.getElementById("logoutButton");
        const feed = document.getElementById("feed");
        const postItem = document.getElementById("separatePostItem");

        if (loginForm) LoginHandler(loginForm);
        if (registerForm) RegisterHandler(registerForm);
        if (userAbout) UserHandler(userAbout);
        if (logoutButton) LogoutHandler(logoutButton);

        if (postItem && !postFetched) {
            postFetched = true;
            PostHandler(postItem)
        }

        // Fetch posts ONLY ONCE
        if (feed && !postsFetched) {
            postsFetched = true; // Mark as fetched
            PostsHandler(feed);
        }
        if (!feed) { // if user goes to another page variable resets
            postsFetched = false;
        }
        if (!postItem) {
            postFetched = false;
        }

        checkAuthNav();
    })
    observer.observe(contentDiv, { childList: true, subtree: true });

    // Checking if user is registered
    async function checkAuthNav() {
        try {
            const response = await fetch("http://localhost:8080/api/user", { credentials: "include" });
            if (!response.ok) {
                if (response.status === 401) {
                    console.warn("Custom Message: You are not logged in.");
                }
                document.getElementById("authorized").style.display = "none";
                document.getElementById("unauthorized").style.display = "block";
            } else if (response.ok) {
                document.getElementById("authorized").style.display = "block";
                document.getElementById("unauthorized").style.display = "none";
            }
        } catch (error) {
            console.error("Auth check failed:", error);
        }
    }
});



