import PostsForm from './views/Posts.js';
import PostForm from './views/Post.js';
import LoginForm from './views/Login.js';
import RegisterForm from './views/Register.js';
import UserAbout from './views/User.js';
import UsersList from "./views/UsersList.js";
import WelcomePage from './views/WelcomePage.js';
import CreatePost from './views/CreatePost.js';

import LoginHandler from './handlers/auth/LoginHandler.js';
import RegisterHandler from './handlers/auth/RegisterHandler.js';
import PostsHandler from './handlers/post/PostsHandler.js';
import PostHandler from './handlers/post/PostHandler.js';
import UserHandler from './handlers/user/UserHandler.js';
import UsersHandler from './handlers/user/UsersHandler.js';
import LogoutHandler from './handlers/auth/LogoutHandler.js';
import CreatePostHandler from "./handlers/post/CreatePostHandler.js";

import FetchComments from './handlers/comment/CommentHandler.js';

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
    window.scrollTo(0, 0);
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
        { path: "/users", view: UsersList },
        { path: "/create-post", view: CreatePost },
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
            document.querySelector("#content").innerHTML = "<h2>404 - Page Not Found</h2>";
            // navigateTo("/posts");
        }, 2000);
        return;
    }

    const view = new match.route.view(getParams(match));

    document.querySelector("#content").innerHTML = await view.getHtml();
};

window.addEventListener("popstate", router);

const fetchedFlags = {
    posts: false,
    post: false,
    users: false,
};

// Массив конфигураций обработчиков
const handlers = [
    { id: "loginForm", handler: LoginHandler },
    { id: "registerForm", handler: RegisterHandler },
    { id: "logoutButton", handler: LogoutHandler },
    { id: "userAbout", handler: UserHandler },
    { id: "postForm", handler: CreatePostHandler },
    { id: "feed", handler: PostsHandler, flag: "posts" },
    { id: "separatePostItem", handler: PostHandler, flag: "post", extra: FetchComments },
    { id: "usersList", handler: UsersHandler, flag: "users" },
];

// Универсальная функция подключения обработчиков
function attachHandler(id, handler, flag, extraFn) {
    const element = document.getElementById(id);

    if (element && (!flag || !fetchedFlags[flag])) {
        if (flag) fetchedFlags[flag] = true;
        handler(element);
        if (extraFn) extraFn();
    } else if (!element && flag) {
        fetchedFlags[flag] = false;
    }
}

document.addEventListener("DOMContentLoaded", () => {
    // Routing
    bindLinkEvents();
    router();

    // Connecting server to client
    const contentDiv = document.getElementById("content");

    const observer = new MutationObserver(() => {
        // Authentication
        handlers.forEach(({ id, handler, flag, extra }) =>
            attachHandler(id, handler, flag, extra)
        );

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



