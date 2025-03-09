import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Posts");
    }

    async getHtml() {
        await checkAuth();
        console.log(IsAuthenticated);
        if (IsAuthenticated) {
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