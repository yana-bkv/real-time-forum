import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Login");
    }

    async getHtml() {
        return `
            <h2>Login</h2>
            <div class="login">
                <form id="loginForm">
                    <label for="emailLogin">Email</label>
                    <input type="email" id="emailLogin" name="email" placeholder="Email" required>
            
                    <label for="passwordLogin">Password</label>
                    <input type="password" id="passwordLogin" name="password" placeholder="Password" required minlength="4">
            
                    <button type="submit">Login</button>
                </form>
            </div>
        `;
    }
}