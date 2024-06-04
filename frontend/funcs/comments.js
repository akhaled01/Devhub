import { BACKENDURL } from "./vars";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import { SetSessionStorageStats } from "./utils";

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

  const commentsDiv = document.getElementById("comments-wrapper"); // Get comments div
  commentsDiv.innerHTML = ""; // Clear comments div you are double rendering
  let data = await fetch_comments(postId); // fetch post comments

  // Render comments
  if (data) {
    data.forEach((comment) => {
      var gender = comment.user.gender;
      if (comment.liked) {
        commentsDiv.innerHTML += `${render_comment_card(
          comment,
          heart,
          gender
        )}`;
      } else {
        commentsDiv.innerHTML += `${render_comment_card(
          comment,
          noheart,
          gender
        )}`;
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
    commentsDiv.innerHTML = `<h4 id="no-comments"> This Post Have No Comments.</h4>`;
  }
}

export const handle_action_like = async (event) => {
  const commentId = event.target.id;
  const likeBtn = event.target;
  // Get like button image

  const likeImgSrc = likeBtn.getAttribute("src");
  const like_counter_div = likeBtn.parentNode.nextElementSibling;
  // Toggle like button image
  await fetch_like_comment_action_API(commentId); // Toggle like or dislike API
  if (likeImgSrc === noheart) {
    // Like comment
    likeBtn.setAttribute("src", heart);
    like_counter_div.textContent = parseInt(like_counter_div.textContent) + 1;
  } else {
    // Unlike comment
    likeBtn.setAttribute("src", noheart);
    like_counter_div.textContent = parseInt(like_counter_div.textContent) - 1;
  }

  await RedoStats();
};

export const RedoStats = async () => {
  await SetSessionStorageStats();
  let Number_of_liked_comments = sessionStorage.getItem(
    "Number_of_liked_comments"
  );
  let Number_of_comments = sessionStorage.getItem("Number_of_comments");
  let Number_of_liked_posts = sessionStorage.getItem("Number_of_liked_posts");
  let Number_of_posts = sessionStorage.getItem("Number_of_posts");
  document.getElementById(
    "UserInfo-div"
  ).innerHTML = `<p class="UserName-p" style="font-size:20px">${sessionStorage.getItem(
    "username"
  )}</p>
        <div class="user-stats" style="font-size: 12px;">
        <p class="user-postd">Posts: ${Number_of_posts}</p>
        <p class="user-likes">Liked Posts: ${Number_of_liked_posts}</p>
        <p class="user-comments">Comments: ${Number_of_comments}</p>
        <p class="user-comments">Liked Comments: ${Number_of_liked_comments}</p>`;
};

export const fetch_like_comment_action_API = async (commentId) => {
  const response = await fetch(`${BACKENDURL}/likeComment/${commentId}`, {
    method: "POST",
    credentials: "include",
  });
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
export const render_comment_card = (comment, like_img, gender) => {
  return /*html*/ `<div class="comment" id="comment_${comment.uuid}">
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
</div>`;
};
