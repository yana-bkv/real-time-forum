import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    loginUser: (data) => apiFetch(`${BASE_URL}/api/login`, {method: "POST", body: data}),
    logoutUser: () => apiFetch(`${BASE_URL}/api/logout`, {method: "POST"}),
    registerUser: (data) => apiFetch(`${BASE_URL}/api/register`, {method: "POST", body: data}),

    getUsers: () => apiFetch(`${BASE_URL}/api/users`),
    getAuthUser: () => apiFetch(`${BASE_URL}/api/user`),
    getUserById: (id) => apiFetch(`${BASE_URL}/api/user/${id}`),
}