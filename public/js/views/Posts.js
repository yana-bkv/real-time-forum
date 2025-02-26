import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Posts");
    }

    async getHtml() {
        return `
          <div id="feed">
            <h1>Posts</h1>
          </div>
        `;
    }
}