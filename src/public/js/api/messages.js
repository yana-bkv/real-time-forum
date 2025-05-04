import {BASE_URL} from "./config.js";
import {apiFetch} from "./apiFetch.js";

export const api = {
    getMsg: (user,peer) => apiFetch(`${BASE_URL}/api/messages/${user}/${peer}`)
}