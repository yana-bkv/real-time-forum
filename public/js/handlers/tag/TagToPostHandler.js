import {api} from '../../api/tag.js';

export async function AssignTagToPostHandler(postId) {
    const category_list = [];

    const input = document.querySelector("input[name='categories']:checked");
    if (input) {
        category_list.push(input.value);
    }

    const data = {
        postId: category_list
    }

    await api.addCategoriesToPostId(postId, data)
}

