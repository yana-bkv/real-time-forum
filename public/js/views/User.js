import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Posts");
    }

    async getHtml() {
        return `
          <div id="userAbout">
            <h1>User</h1>
            <p class="userName"></p>
            <p class="userEmail"></p>
          </div>
        `;
    }
}