import { BACKENDURL } from "./vars";
// Fetch the comments data
export async function fetchComments(postId) {
    if (postId === null) {
        console.error('postId is null');
        return;
    }
    console.log(postId);
    const response = await fetch(`${BACKENDURL}/comments/${postId}`, {
        credentials: "include",
    });
    const data = await response.json();
    console.log(data, data.comments);

    const commentsDiv = document.getElementById("comments");
    if (data) {
        data.forEach((comment) => {
            commentsDiv.innerHTML += `
              <div class="comment">
                  <div class="comment-header">
                      <div class="c-profileInfo">
                          <div class="c-profile-pic"></div>
                          <div class="c-nickname">${comment.user.username}</div>
                      </div>
                      <div class="c-creationDate">${new Date(
                        comment.creationDate
                    ).toDateString()}</div>
                  </div>
                  <div class="p-main">${comment.content}</div>
              </div>
          `;
        });
    } else {
        console.error('No comments data received');
        commentsDiv.innerHTML = `<h4 style="color:white;"> This Post Have No Comments. </h4>`
    }
}
