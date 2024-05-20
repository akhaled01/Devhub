import { BACKENDURL } from "./vars";
import noheart from "../assets/unliked.svg";
import comment from "../assets/comment.svg";
import heart from "../assets/liked.svg";
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
    credentials: "include",
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
  document.getElementById("main_wrapper").innerHTML = "";
  posts_in_json.forEach((post_data) => {
    var gender = post_data.user.gender;
    let liked_img = noheart;
    if (post_data.liked) {
      liked_img = heart;
    }
    let text = post_data.content + "";

    if (text.length > 255) {
      text = text.slice(0, 255 - "...".length) + "..."; // turncate
    }

    document.getElementById("main_wrapper").innerHTML += `<a href="/post/${
      post_data.id
    }"><div class="f-post ${!post_data.Image_Path ? "noimage" : ""}" id=${
      post_data.id
    }>
  <div class="p-header">
    <div class="p-profileInfo">
      <div class="p-profile-pic gender-${post_data.user.gender}">${
      post_data.user.username[0].toUpperCase()
    }</div>
      <div class="p-nickname">${post_data.user.username}</div>
      </div>
      <div class="p-creationDate">${new Date(
        post_data.creationDate
      ).toDateString()}</div>
        </div>
        <div class="p-main">
        <div class="p-content">
        ${text}
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
        <img src="${liked_img}" alt="like" />
        </div>
        <div class="p-likeStat">${post_data.likes}</div>
        </div>
        <div class="p-commentCount">
        <img src="${comment}" alt="comment" />
        <div class="p-comment-Stat">${post_data.number_of_comments}</div>
        </div>
        </div>
        </div>
        <div class="p-Category">
        <p>#${post_data.category}</p>
        </div>
</div>
</a>
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
    const default_profile_pic = "../assets/defaultPfp.svg";
    fetch(default_profile_pic)
      .then((response) => response.blob())
      .then((blob) => {
        const reader = new FileReader();
        reader.onload = function (e) {
          const bs64str = e.target.result;
          console.log("Default Base64 Image:", bs64str);
          callback(bs64str); // Call the callback function with the base64 string of the custom image
        };
        reader.readAsDataURL(blob);
      })
      .catch((error) => console.error("Error fetching custom file:", error));
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
  sessionStorage.setItem("gender", json_data.gender);
};

export const SetSessionStorageStats = async () => {
  const res = await fetch(BACKENDURL + "/userstats", {
    credentials: "include",
  });
  if (res.ok) {
    const json_data = await res.json();
    console.log("Fetching user stats",json_data);
    sessionStorage.setItem("Number_of_liked_comments", json_data.Number_of_liked_comments);
    sessionStorage.setItem("Number_of_liked_posts", json_data.Number_of_liked_posts);
    sessionStorage.setItem("Number_of_comments", json_data.Number_of_comments);
    sessionStorage.setItem("Number_of_posts", json_data.Number_of_posts);
  } else {
    console.error("Error fetching user stats");
  }
};

/**
 * Renders a new chat message to the message area
 * @param {string} message - the message content
 * @param {boolean} is_self -
 * @param {*} name - name of send (if `is_self` is false)
 * @param {Date} time - time of message
 */
export const NewChatMessage = (
  message,
  is_self,
  name = "",
  time = new Date()
) => {
  const messageElement = document.createElement("div");
  const actualMessage = document.createElement("div");
  // is_self checks if the message came from the current user, not the
  // other one
  if (is_self) {
    messageElement.classList.add("mself");
    actualMessage.classList.add("self");
    actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">You</div>
              <div class="date">${time.toDateString()}</div>
            </div>`;
  } else {
    messageElement.classList.add("m");
    actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">${name}</div>
              <div class="date">${time.toDateString()}</div>
            </div>`;
  }

  actualMessage.classList.add("message");

  const content = document.createElement("p");
  const chatArea = document.getElementById("message_space");
  content.textContent = message;
  actualMessage.appendChild(content);
  messageElement.appendChild(actualMessage);
  chatArea.appendChild(messageElement);
  chatArea.scrollTop = chatArea.scrollHeight; // Scroll to bottom
};

export function convertImageToBase64(file) {
  if (!file) {
    return null;
  }

  const reader = new FileReader();

  return new Promise((resolve, reject) => {
    reader.onloadend = function () {
      if (reader.readyState === FileReader.DONE) {
        const base64String = reader.result;
        console.log(base64String);
        resolve(base64String);
      } else {
        reject(new Error("Error reading file"));
      }
    };

    reader.readAsDataURL(file);
  });
}

/**
 * Sorts online user as per required
 * @param {*} arr
 * @returns
 */
export const sortByOnlineAndName = (arr) => {
  return arr.req_Content.sort((a, b) => {
    return a.username.localeCompare(b.username);
  });
};
