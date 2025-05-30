import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    createTag: (data) => apiFetch(`${BASE_URL}/api/category`, {method: "POST", body: data}),
    getTags: () => apiFetch(`${BASE_URL}/api/categories`),
    getTagById: (id) => apiFetch(`${BASE_URL}/api/category/${id}`),
    deleteTagById: (id) => apiFetch(`${BASE_URL}/api/category/${id}`, {method: "DELETE"}),

    addCategoriesToPostId:(id, data) => apiFetch(`${BASE_URL}/api/post/${id}/category`, {method: "POST", body: data}),
    getTagsByPostId: (id) => apiFetch(`${BASE_URL}/api/post/${id}/categories`),
}