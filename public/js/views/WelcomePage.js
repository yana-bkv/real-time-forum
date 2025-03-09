import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated, AuthenticatedUserName } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Posts");
    }

    async getHtml() {
        await checkAuth();
        if (IsAuthenticated) {
            return `
          <h2>
            Hi ${AuthenticatedUserName}. Welcome back.
          </h2>
        `;
        } else {
            return `
                <h2>Welcome!</h2>
            `;
        }
    }
}