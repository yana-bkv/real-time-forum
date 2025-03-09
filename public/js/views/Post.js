import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Post");
    }

    async getHtml() {
        await checkAuth();
        if (IsAuthenticated) {
            return `
          <div id="postTitle">
            <div id="postBody"></div>
            <p class="postAuthor"></p>
            <p class="postTime"></p>
          </div>
        `;
        } else {
            return `
                <h2>User need to be logged in to view profile</h2>
            `;
        }
    }
}