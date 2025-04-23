import AbstractView from "./AbstractView.js";
import { authService } from "../helpers/AuthServiceCheck.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Post");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();
        if (isAuthenticated) {
            return `
                <div class="post-block">
                <h1>Post</h1>
                <div id="separatePostItem"></div>
                <button id="deletePost">Delete post</button>
                
                <div id="comments-container"></div>
                </div>`
        } else {
            return `
                <h2>User is not authenticated to view posts</h2>
            `;
        }
    }
}