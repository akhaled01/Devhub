import { BACKENDURL } from "./vars";
import noheart from "../assets/unliked.svg";
import comment from "../assets/comment.svg";
import heart from "../assets/liked.svg";
import { PostModal } from "../pages";
import { ws } from "../main";
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
  const mainWrapper = document.getElementById("main_wrapper");
  mainWrapper.innerHTML = "";

  posts_in_json.forEach((post_data) => {
    const gender = post_data.user.gender;
    const liked_img = post_data.liked ? heart : noheart;
    let text_html = ``;
    let text = post_data.content + "";

    // console.log(text);

    if (text.length > 100) {
      text = text.slice(0, 103 - "...".length) + "..."; // truncate
    }

    text.split("@").forEach((str) => {
      text_html += str + `<br />`;
    });

    const postHTML = `
      <div class="f-post ${!post_data.Image_Path ? " noimage" : ""}" id="post_${
      post_data.id
    }">
        <div class="p-header">
          <div class="p-profileInfo">
            <div class="p-profile-pic gender-${gender}">${post_data.user.username[0].toUpperCase()}</div>
            <div class="p-nickname">${post_data.user.username}</div>
          </div>
          <div class="p-creationDate">${new Date(
            post_data.creationDate
          ).toLocaleDateString()}</div>
        </div>
        <div class="p-main">
          <div class="p-content">
            ${text_html}
            ${
              post_data.Image_Path
                ? `<div class="p-image"><img src=${post_data.Image_Path} alt="post image"></div>`
                : ""
            }
          </div>
          <div class="p-stats">
            <div class="p-likeCount">
              <div class="p-likeBtn"><img src="${liked_img}" alt="like" /></div>
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
    `;

    mainWrapper.innerHTML += postHTML;
  });

  posts_in_json.forEach((post_data) => {
    const postElement = document.getElementById(`post_${post_data.id}`);
    if (postElement) {
      postElement.addEventListener("click", (e) => {
        PostModal(post_data.id);
      });
    }
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
    sessionStorage.setItem("user_token", json_data.session_id);
    sessionStorage.setItem("username", json_data.username);
    sessionStorage.setItem("email", json_data.email);
    sessionStorage.setItem("avatar", json_data.encoded_avatar);
    sessionStorage.setItem("gender", json_data.gender);
    sessionStorage.setItem(
      "Number_of_liked_comments",
      json_data.Number_of_liked_comments
    );
    sessionStorage.setItem(
      "Number_of_liked_posts",
      json_data.Number_of_liked_posts
    );
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
              <div class="date">${time.toLocaleDateString()}</div>
            </div>`;
  } else {
    messageElement.classList.add("m");
    actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">${name}</div>
              <div class="date">${time.toLocaleDateString()}</div>
            </div>`;
  }

  actualMessage.classList.add("message");

  const content = document.createElement("p");
  const chatArea = document.getElementById("message_space");
  content.textContent = message;
  actualMessage.appendChild(content);
  messageElement.appendChild(actualMessage);
  if (chatArea) {
    chatArea.appendChild(messageElement);
    chatArea.scrollTop = chatArea.scrollHeight; // Scroll to bottom
  }
};

/**
 * Ensure correct pagination of historical messages
 */
export const PaginateHistoricalMessage = (
  message,
  is_self,
  name = "",
  time = new Date()
) => {
  const chatArea = document.getElementById("message_space");
  const messageElement = document.createElement("div");
  const actualMessage = document.createElement("div");
  let prev_scroll_pos = chatArea.scrollTop;
  // is_self checks if the message came from the current user, not the
  // other one
  if (is_self) {
    messageElement.classList.add("mself");
    actualMessage.classList.add("self");
    actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">You</div>
              <div class="date">${time.toLocaleDateString()}</div>
            </div>`;
  } else {
    messageElement.classList.add("m");
    actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">${name}</div>
              <div class="date">${time.toLocaleDateString()}</div>
            </div>`;
  }

  const content = document.createElement("p");
  content.textContent = message;
  actualMessage.appendChild(content);

  actualMessage.classList.add("message");

  messageElement.appendChild(actualMessage);

  chatArea.insertAdjacentElement("afterbegin", messageElement);
  chatArea.scrollTop = prev_scroll_pos;
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

/**
 *
 * Marks typing in progress
 *
 * @param {string} username - username of the typer
 * @param {bool} is_typing - status of typing
 */
export const MarkTyping = (is_typing = false) => {
  const profileDiv = document.getElementById("r-profile");

  if (is_typing) {
    const tipDiv = document.createElement("p");
    tipDiv.id = "typing_in_progress";
    tipDiv.innerText = "Typing..";

    profileDiv.append(tipDiv);
  } else {
    profileDiv.querySelector("#typing_in_progress").remove();
  }
};

export function delay(time) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

export const throttle = (func, limit) => {
  let lastFunc;
  let lastRan;
  return function (...args) {
    const context = this;
    if (!lastRan) {
      func.apply(context, args);
      lastRan = Date.now();
    } else {
      clearTimeout(lastFunc);
      lastFunc = setTimeout(function () {
        if (Date.now() - lastRan >= limit) {
          console.log(lastRan);
          func.apply(context, args);
          lastRan = Date.now();
        }
      }, limit - (Date.now() - lastRan));
    }
  };
};

export const sendTypingStart = throttle(() => {
  console.log("send");
  ws.send(
    JSON.stringify({
      type: "typing-event",
      req_Content: {
        recipient_name: sessionStorage.getItem("chat_partner"),
        signal_type: "start",
        Sender_name: sessionStorage.getItem("username"),
      },
    })
  );
}, 1000);

// Throttled function to send the "stop" typing event
export const sendTypingStop = throttle(() => {
  console.log("stop");
  ws.send(
    JSON.stringify({
      type: "typing-event",
      req_Content: {
        recipient_name: sessionStorage.getItem("chat_partner"),
        signal_type: "stop",
        Sender_name: sessionStorage.getItem("username"),
      },
    })
  );
}, 2000);
