import LoginHandler from './handlers/LoginHandler.js';
import RegisterHandler from './handlers/RegisterHandler.js';
import UserHandler from './handlers/UserHandler.js';
import LogoutHandler from './handlers/LogoutHandler.js';
import CreatePostHandler from './handlers/CreatePostHandler.js';
import PostsHandler from './handlers/PostsHandler.js';
import PostHandler from './handlers/PostHandler.js';
import UsersHandler from './handlers/UsersHandler.js';
import FetchComments from './comments/FetchComments.js';
import { checkAuthNav } from './auth/isAuthenticatedUser.js';

const fetchedFlags = {
    posts: false,
    post: false,
    users: false
};

export function initHandlers() {
    const loginForm = document.getElementById("loginForm");
    const registerForm = document.getElementById("registerForm");
    const logoutButton = document.getElementById("logoutButton");
    const userAbout = document.getElementById("userAbout");
    const usersList = document.getElementById("usersList");
    const feed = document.getElementById("feed");
    const postItem = document.getElementById("separatePostItem");
    const createPostForm = document.getElementById("postForm");

    if (loginForm) LoginHandler(loginForm);
    if (registerForm) RegisterHandler(registerForm);
    if (userAbout) UserHandler(userAbout);
    if (logoutButton) LogoutHandler(logoutButton);
    if (createPostForm) CreatePostHandler(createPostForm);

    if (feed && !fetchedFlags.posts) {
        fetchedFlags.posts = true;
        PostsHandler(feed);
    }
    if (!feed) fetchedFlags.posts = false;

    if (postItem && !fetchedFlags.post) {
        fetchedFlags.post = true;
        PostHandler(postItem);
        FetchComments();
    }
    if (!postItem) fetchedFlags.post = false;

    if (usersList && !fetchedFlags.users) {
        fetchedFlags.users = true;
        UsersHandler(usersList);
    }
    if (!usersList) fetchedFlags.users = false;

    checkAuthNav();
}