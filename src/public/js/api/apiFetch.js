export async function apiFetch(url, options = {}) {
    const {
        method = "GET",
        headers = {},
        body = null,
        withAuth = true
    } = options;

    const fetchOptions = {
        method,
        headers: {
            "Content-Type": "application/json",
            ...headers,
        },
        credentials: withAuth ? "include" : "same-origin",
    };

    if (body) {
        fetchOptions.body = JSON.stringify(body);
    }

    try {
        const response = await fetch(url, fetchOptions);
        const contentType = response.headers.get("Content-Type");

        let result;

        if (contentType && contentType.includes("application/json")) {
            result = await response.json();
        } else {
            result = await response.text(); // ← если не JSON, хотя это редко
        }

        if (!response.ok) {
            throw new Error(result.message || "Request failed.");
        }

        return result;
    } catch (error) {
        console.error("API Error:", error.message);
        throw error;
    }
}

export const BASE_URL = "http://localhost:8080";
