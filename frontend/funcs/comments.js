import { BACKENDURL } from "./vars";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";

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
    console.log(data);

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
                  <div class="p-stats">
                    <div class="p-likeCount">
                      <div class="p-likeBtn">
                          <img src="${noheart}" alt="like"/>
                      </div>
                      <div class="p-likeStat">${comment.likes}</div>
                    </div>
                  </div>
              </div>
          `;
            const likeImages = document.querySelectorAll(".p-likeBtn img");

            console.log(likeImages);

            likeImages.forEach((likeBtn) => {
                console.log(likeBtn.getAttribute("src"));

                likeBtn.addEventListener("click", () => {
                    if (likeBtn.getAttribute("src") === noheart) {
                        likeBtn.setAttribute("src", heart);
                        console.log("liked");
                        // add other like event
                    } else {
                        likeBtn.setAttribute("src", noheart);
                        console.log("unliked");
                        // add other unlike event
                    }

                    try {
                        async function likeComment(commentID) {
                            const res = await fetch(BACKENDURL + `/likeComment/` + commentID, {
                                method: "POST",
                                body: JSON.stringify({
                                    Comment_id: commentID,
                                    // User_id: userId,
                                }),
                                credentials: "include",
                                //   headers: {
                                //       "Content-Type": "application/json",
                                //   },
                            });
                            console.log(response);
                            if (response.ok) {
                                // const updatedComment = await response.json();
                                // Update the comment data in the UI
                                const commentElement = document.querySelector(`[data-comment-id="${commentID}"]`);
                                if (commentElement) {
                                    commentElement.querySelector('.comment-likes').textContent = respone.json.likes;
                                    commentElement.querySelector('.like-icon').classList.toggle('liked', response.json.liked);
                                }
                            } else {
                                console.error('Error creating the like:', response.statusText);
                            }
                            //   const data = await res.json();
                            return res;
                        }

                        async function likeCommentWrapper(comment) {
                            console.log(comment.uuid);

                            const res = await likeComment(comment.uuid);

                            if (res.status === 200) {
                                // window.location.reload();
                            } else {
                                throw new Error(res.status, res.statusText);
                            }
                        }

                        likeCommentWrapper(comment);
                    } catch (error) {
                        console.error('Error creating the like2:', error);
                        alert('Failed to create comment like. Please try again later.');
                    }
                });
            });
        });
    } else {
        console.error('No comments data received');
        commentsDiv.innerHTML = `<h4 style="color:white;"> This Post Have No Comments. </h4>`
    }
}
