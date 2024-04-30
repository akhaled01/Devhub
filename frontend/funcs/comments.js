// Fetch the comments data
export async function fetchComments() {
    const response = await fetch(`/api/comments?Post_id=${postId}`);
    const data = await response.json();
  
    const commentsDiv = document.getElementById("comments");
    data.comments.forEach((comment) => {
      commentsDiv.innerHTML += `
              <div class="comment">
                  <div class="comment-header">
                      <div class="c-profileInfo">
                          <div class="c-profile-pic"></div>
                          <div class="c-nickname">${comment.author}</div>
                      </div>
                      <div class="c-creationDate">${comment.creationDate}</div>
                  </div>
                  <div class="p-main">${comment.content}</div>
              </div>
          `;
    });
  }
  
  