import AbstractView from "./AbstractView.js";
import { authService } from "../helpers/AuthServiceCheck.js";


export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Create Post");
    }

    async getHtml() {
        const isAuthenticated = await authService.checkAuth();

        if (isAuthenticated) {
            return `
                <h2>Create Post</h2>
                <div class="postFormCreate">
                    <form id="postForm">
                        <label for="postTitle">Post Title</label>
                        <input type="text" id="postTitle" name="title" placeholder="Title" required>
                
                        <label for="postBody">Main text</label>
                        <textarea id="postBody" name="postBody" placeholder="Text..." rows="5"></textarea>
                        
                        <div class="postCategory"></div>
                        
                        <button type="submit">Create</button>
                    </form>
                </div>
            `;
        } else {
            return `
                <h2>User is not authenticated to view post</h2>
            `;
        }
    }
}
