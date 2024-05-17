import { BACKENDURL } from "./vars";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";

// Fetching from API
const fetch_comments = async (postId) => {
  const response = await fetch(`${BACKENDURL}/comments/${postId}`, {
    credentials: "include",
  });
  const data = await response.json();
  return data;
};

// Render comments
export async function render_comments(postId) {
  if (postId === null) {
    console.error("postId is null");
    return;
  }

  const commentsDiv = document.getElementById("comments"); // Get comments div

  let data = await fetch_comments(postId);  // fetch post comments

  // Render comments
  if (data) {
    data.forEach((comment) => {
      var gender = comment.user.gender;
      if (comment.liked) {
        commentsDiv.innerHTML += `${render_comment_card(comment, heart,gender)}`;
      } else {
        commentsDiv.innerHTML += `${render_comment_card(comment, noheart,gender)}`;
      }

      // Like button click event
      const Like_Count_divs = document.querySelectorAll(".p-likeCountC");
      Like_Count_divs.forEach((Like_Count_div) => {
        // Handle hearts
        const likeBtn = Like_Count_div.querySelector(".p-likeBtnC");
        likeBtn.addEventListener("click", handle_action_like);

        const likeCount = Like_Count_div.querySelector(".p-likeStat");

      });
    });
  } else {
    console.error("No comments data received");
    commentsDiv.innerHTML = `<h4 style="color:white;"> This Post Have No Comments. </h4>`;
  }
}

export const handle_action_like = async (event) => {
  const commentId = event.target.id;
  const likeBtn = event.target;
  // Get like button image

  const likeImgSrc = likeBtn.getAttribute("src");
  const like_counter_div = likeBtn.parentNode.nextElementSibling
  // Toggle like button image
  await fetch_like_comment_action_API(commentId) // Toggle like or dislike API
  if (likeImgSrc === noheart) {
    // Like comment
    likeBtn.setAttribute("src", heart);
    like_counter_div.textContent = parseInt(like_counter_div.textContent) + 1;
  } else {
    // Unlike comment
    likeBtn.setAttribute("src", noheart);
    like_counter_div.textContent = parseInt(like_counter_div.textContent) - 1;
  }
};

export const fetch_like_comment_action_API = async (commentId) => {
  const response = await fetch(
    `${BACKENDURL}/likeComment/${commentId}`,
    {
      method: "POST",
      credentials: "include",
    }
  );
  const data = await response.json();
  return data;
};



/**
 *
 * @param {*} comment
 * @param {SVGAElement} like_img
 * @returns
 */
/* Render comment card */
export const render_comment_card = (comment, like_img,gender) => {
  return /*html*/`<div class="comment" id="${comment.uuid}">
  <div class="comment-header">
      <div class="c-profileInfo">
          <div class="c-profile-pic gender-${gender}"></div>
          <div class="c-nickname">${comment.user.username}</div>
      </div>
      <div class="c-creationDate">${new Date(
    comment.creationDate
  ).toDateString()}</div>
  </div>
  <div class="p-main">${comment.content}</div>
  <div class="p-stats">
    <div class="p-likeCountC">
      <div class="p-likeBtnC">
          <img src="${like_img}" alt="like" id="${comment.uuid}"/>
      </div>
      <div class="p-likeStat">${comment.likes}</div>
    </div>
  </div>
</div>`

}
