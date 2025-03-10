import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("User Cabinet");
    }

    async getHtml() {
        await checkAuth();
        if (IsAuthenticated) {
            return `
          <div id="userAbout">
            <h1>Your Cabinet</h1>
            <p class="userName"></p>
            <p class="userEmail"></p>
          </div>
        `;
        } else {
            return `
                <h2>User need to be logged in to view profile</h2>
            `;
        }
    }
}