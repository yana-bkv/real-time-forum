import AbstractView from "./AbstractView.js";
import { checkAuth, IsAuthenticated, AuthenticatedUserName } from "./isAuthenticatedUser.js";


export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Login");
    }

     async getHtml() {
        await checkAuth();
        if (!IsAuthenticated) {
            return `
                <h2>Login</h2>
                <div class="login">
                    <form id="loginForm">
                        <label for="emailLogin">Email or Username</label>
                        <input type="text" id="emailLogin" name="email" placeholder="Email or Username" required>
                
                        <label for="passwordLogin">Password</label>
                        <input type="password" id="passwordLogin" name="password" placeholder="Password" required minlength="4">
                
                        <button type="submit">Login</button>
                    </form>
                </div>
            `;
        } else {
            return `
            <h2>Hi ${AuthenticatedUserName}, you are succefully logged in</h2>
            `;
        }
    }
}
