import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated } from "./isAuthenticatedUser.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Register");
    }

    async getHtml() {
        await checkAuth();
        if (!IsAuthenticated) {
        return `
            <h2>Register</h2>
            <div class="register">
                <form id="registerForm">
                    <label for="username">Username</label>
                    <input type="text" id="username" name="username" placeholder="Username" required>
            
                    <label for="email">Email</label>
                    <input type="email" id="email" name="email" placeholder="Email" required>
            
                    <label for="password">Password</label>
                    <input type="password" id="password" name="password" placeholder="Password" required minlength="4">
            
                    <button type="submit">Register</button>
                </form>
            </div>
        `;} else {
            return `
            <h2>User already authenticated</h2>
            `;
        }
    }
}