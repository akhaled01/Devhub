import { BACKENDURL, post_wrapper } from "./vars";
import noheart from "../assets/unliked.svg";
import comment from "../assets/comment.svg";

/**
 *
 * Follow up login after signup
 *
 * @param {string} email
 * @param {string} pass
 */
export const Flogin = async (email, pass) => {
  if (!email || !pass) {
    return;
  }

  const res = await fetch(BACKENDURL + "/auth/login", {
    method: "POST",
    body: JSON.stringify({
      credential: email,
      password: pass,
    }),
  });

  if (res.ok) {
    SetSessionStorage(await res.json());
    window.location.assign("/");
  }
};

/**
 * Function to Update CSS en routing
 * @param {Stylesheet} stylesheet - Path to css file
 */
export const UpdateCSS = (stylesheet) => {
  const linkElement = document.getElementById("page-styles");
  if (linkElement) {
    linkElement.href = stylesheet;
    const styleTags = document.querySelectorAll("style");
    styleTags.forEach((tag) => tag.remove());
  } else {
    console.error("Page stylesheet link not found");
  }
};

/**
 * Takes in an array of json posts and renders them
 * on the index page
 * @param {any[]} posts_in_json
 */
export const AssemblePosts = (posts_in_json = []) => {
  posts_in_json.forEach((post_data) => {
    post_wrapper.innerHTML += `<div class="f-post ${
      !post_data.Image_Path ? "noimage" : ""
    }" id=${post_data.id}>
  <div class="p-header">
    <div class="p-profileInfo">
      <div class="p-profile-pic"></div>
      <div class="p-nickname">${post_data.user.username}</div>
    </div>
    <div class="p-creationDate">${post_data.CreationDate}</div>
  </div>
  <div class="p-main">
    <div class="p-content">
      ${post_data.content}
      ${
        post_data.Image_Path
          ? `<div class="p-image">
        <img src=${post_data.Image_Path} alt="post image">
      </div>`
          : ""
      }
    </div>
    <div class="p-stats">
      <div class="p-likeCount">
        <div class="p-likeBtn">
          <img src="${noheart}" alt="like" />
        </div>
        <div class="p-likeStat">${post_data.likes}</div>
      </div>
      <div class="p-commentCount">
        <img src="${comment}" alt="comment" />
        <div class="p-comment-Stat">${post_data.number_of_comments}</div>
      </div>
    </div>
  </div>
</div>
    `;
  });
};

/**
 * Function that encodes the avatar uploaded by the
 * user. This func is async, so it takes in a callback
 * Instead
 */
export const EncodeBase64Image = (callback) => {
  const fileInput = document.getElementById("avatar");

  if (fileInput.files.length > 0) {
    const file = fileInput.files[0];
    const reader = new FileReader();

    reader.onload = function (e) {
      const bs64str = e.target.result;
      console.log("Base64 Image:", bs64str);
      callback(bs64str); // Call the callback function with the base64 string
    };

    reader.readAsDataURL(file);
  } else {
    console.log("Please select an image file.");
  }
};

/**
 * Sets the required sessionStorage Params
 * @param {*} json_data
 */
export const SetSessionStorage = (json_data) => {
  sessionStorage.setItem("user_token", json_data.session_id);
  sessionStorage.setItem("username", json_data.username);
  sessionStorage.setItem("email", json_data.email);
  sessionStorage.setItem("avatar", json_data.encoded_avatar);
};
