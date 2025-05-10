import {api} from '../../api/likes.js';

export async function LikeHandler(postId) {
    const likeSection = document.createElement('div');

    const pLikeBtn = document.createElement('p');
    const likeBtn = document.createElement('button');

    const likeCounter = document.createElement('p');

    let hasLiked = await api.hasLiked(postId);
    let likeCount = await api.getLikeCount(postId);

    likeBtn.textContent = hasLiked.hasLiked ? 'Unlike' : 'Like';
    likeCounter.textContent = `Likes: ${likeCount}`;

    likeBtn.addEventListener('click', async (e) => {
        e.preventDefault();
        if (hasLiked.hasLiked) {
            await api.removeLike(postId);
        } else {
            await api.addLike(postId);
        }
        likeCount = await api.getLikeCount(postId);
        hasLiked.hasLiked = !hasLiked.hasLiked;
        likeBtn.textContent = hasLiked.hasLiked ? 'Unlike' : 'Like';
        likeCounter.textContent = `Likes: ${likeCount}`;
    });

    pLikeBtn.appendChild(likeBtn);
    likeSection.appendChild(pLikeBtn);
    likeSection.appendChild(likeCounter);
    return likeSection;
}