export let IsAuthenticated = false;
export let AuthenticatedUserName;

export async function checkAuth() {
    try {
        const response = await fetch("http://localhost:8080/api/user", {credentials: "include"});

        if (response.status === 401) {
            IsAuthenticated = false;
        } else if (response.ok) {
            IsAuthenticated = true;
            const result = await response.json();
            AuthenticatedUserName = result.username  // Update authentication state
        } else {
            IsAuthenticated = false;
        }

        return IsAuthenticated; // Return the authentication state
    } catch (error) {
        console.error("Auth check failed:", error);
        IsAuthenticated = false; // Assume not authenticated on error
        return false;
    }
}