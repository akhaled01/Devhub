
// Replace with the actual post ID
const postId = 1;

// for later, it should be connected to the back-end
// Fetch the post data
export async function fetchPost() {
  const response = await fetch(`/api/post/${postId}`);
  const data = await response.json();

  const postDiv = document.getElementById("post");
  postDiv.innerHTML = `
        <div class="f-post">
            <div class="p-header">
                <div class="p-profileInfo">
                    <div class="p-profile-pic"></div>
                    <div class="p-nickname">${data.author}</div>
                </div>
                <div class="p-creationDate">${data.creationDate}</div>
            </div>
            <div class="p-main">
                <div class="p-content">
                    ${data.content}
                    <div class="p-image">
                        <img src="${data.image}" alt="post image">
                    </div>
                </div>
                <div class="p-stats">
                    <div class="p-likeCount">
                        <div class="p-likeBtn">
                            <img src="${noheart}" alt="like" />
                        </div>
                        <div class="p-likeStat">${data.likes}</div>
                    </div>
                    <div class="p-commentCount">
                        <img src="${comment}" alt="comment" />
                        <div class="p-comment-Stat">${data.comments}</div>
                    </div>
                </div>
            </div>
        </div>
    `;
}
