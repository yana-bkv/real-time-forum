import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Users List");
    }

    async getHtml() {
        await checkAuth();
        if (IsAuthenticated) {
            return `
          <div id="usersListPage">
            <h1>List of all users</h1>
            <div id="usersList"></div>
          </div>
        `;
        } else {
            return `
                <h2>User need to be logged in to view profile</h2>
            `;
        }
    }
}