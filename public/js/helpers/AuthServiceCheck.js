import {api} from "../api/users.js";

class AuthService {
    constructor() {
        this.isAuthenticated = false;
        this.username = null;
        this.checked = false; // Prevent multiple auth checks
    }

    // Mimic checkAuth() method as if it were defined by an interface
    async checkAuth() {
        if (this.checked) {
            return this.isAuthenticated;
        }

        try {
            const result = await api.getAuthUser();  // Simulated API call
            this.isAuthenticated = true;
            this.username = result.username;
        } catch (error) {
            this.isAuthenticated = false;
            this.username = null;
        }

        this.checked = true;
        console.log("AuthService check auth ", this.username || 'undefined');
        return this.isAuthenticated;
    }

    getUsername() {
        return this.username;
    }
}

export const authService = new AuthService();