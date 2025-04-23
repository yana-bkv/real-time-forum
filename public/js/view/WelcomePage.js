import AbstractView from "./AbstractView.js";
import { authService } from "../helpers/AuthServiceCheck.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Posts");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();
        const authenticatedUserName = await authService.getUsername();
        if (isAuthenticated) {
            return `
          <h2>
            Hi ${authenticatedUserName}. Welcome back.
          </h2>
        `;
        } else {
            return `
                <h2>Welcome!</h2>
            `;
        }
    }
}