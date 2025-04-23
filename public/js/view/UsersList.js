import AbstractView from "./AbstractView.js";
import { authService } from "../helpers/AuthServiceCheck.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Users List");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();
        if (isAuthenticated) {
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