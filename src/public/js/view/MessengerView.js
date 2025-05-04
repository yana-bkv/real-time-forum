import AbstractView from "./AbstractView.js";
import {authService} from "../helpers/AuthServiceCheck.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Post");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();
        if (isAuthenticated) {
            return `
                <div id="messengerBody">
                <div id="messengerHistory"></div>
                <div id="log"></div>
                <form id="messengerForm">
                    <input type="text" id="msg" size="64" autofocus />
                    <input type="submit" value="Send" />
                </form>
                </div>`
        } else {
            return `
                <h2>User is not authenticated to view posts</h2>
            `;
        }
    }
}