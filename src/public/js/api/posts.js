import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    createPost: (data) => apiFetch(`${BASE_URL}/api/post`, {method: "POST", body: data}),
    getPosts: () => apiFetch(`${BASE_URL}/api/posts`),
    getPostById: (id) => apiFetch(`${BASE_URL}/api/post/${id}`),
    deletePostById: (id) => apiFetch(`${BASE_URL}/api/post/${id}`, {method: "DELETE"}),

    getComments: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/comments`),
}