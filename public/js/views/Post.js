import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Post");
    }

    async getHtml() {
        await checkAuth();
        console.log(IsAuthenticated);
        if (IsAuthenticated) {
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