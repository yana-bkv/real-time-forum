import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated, AuthenticatedUserName } from "./isAuthenticatedUser.js";


export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Create Post");
    }

    async getHtml() {
        await checkAuth();
        if (IsAuthenticated) {
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
