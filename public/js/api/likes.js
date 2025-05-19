import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    addLike: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/like`, {method: 'POST'}),
    getLike: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/likes`),
    hasLiked: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/hasLiked`),
    removeLike: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/like`, {method: 'DELETE'}),
    getLikeCount: (postId) => apiFetch(`${BASE_URL}/api/post/${postId}/likeCount`),
}