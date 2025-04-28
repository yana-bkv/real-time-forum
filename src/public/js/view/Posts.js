import AbstractView from "./AbstractView.js";
import {authService} from "../helpers/AuthServiceCheck.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Posts");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();

        if (isAuthenticated) {
            return `
                <div class="posts-block">
                <h1>Posts</h1>
                <div id="feed"></div></div>`
        } else {
            return `
                <h2>User is not authenticated to view posts</h2>
            `;
        }
    }
}