import {api} from '../../api/comments.js';

export function CreateCommentSection(postId) {
    const section = document.createElement('div');
    section.id = "comment-section";

    section.innerHTML = `
        <form id="comment-form">
            <input id="comment" type="text" name="comment" value="" required/>
            <button type="submit">Send</button>
        </form>
    `;

    const commentInput = section.querySelector('#comment');
    const commentForm =  section.querySelector('#comment-form');

    commentForm.addEventListener('submit', async(e) => {
        e.preventDefault();

        const body = commentInput.value.trim();
        if (!body) return;

        const data = {body};

        try {
            await api.createComment(postId, data);
            commentInput.value = '';
            window.location.reload();
        } catch {
            console.error("Error creating comment");
        }
    })

    return section;
}