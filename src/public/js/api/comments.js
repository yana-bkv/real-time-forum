import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    createComment: (postId, data) => apiFetch(`${BASE_URL}/api/post/${postId}/comment`, {method: "POST", body: data}),
    getComments: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/comments`),
    deleteComment: (postId, commentId) => apiFetch(`${BASE_URL}/api/post/${postId}/comment/${commentId}`, {method: "DELETE"})
}