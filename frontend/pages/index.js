import { LoadNav } from "../funcs/navbar";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import imgupload from "../assets/imageupload.svg";
import hashtag from "../assets/hashtag.svg";
import comment from "../assets/comment.svg";
import { OrgIndexPosts } from "../funcs/posts";
import { BACKENDURL } from "../funcs/vars";
import { convertImageToBase64 } from "../funcs/utils";
import { ws } from "../main";
import { SetSessionStorageStats } from "../funcs/utils";
import { creat_comment, fetch_post } from "./post";
import { RedoStats, render_comments } from "../funcs/comments";
import { RecordPostLikeEvent } from "../funcs/likes";

let Number_of_liked_comments = 0;
let Number_of_comments = 0;
let Number_of_liked_posts = 0;
let Number_of_posts = 0;

export const GetSessionStorageStats = async () => {
  await SetSessionStorageStats();
  Number_of_liked_comments = sessionStorage.getItem("Number_of_liked_comments");
  Number_of_comments = sessionStorage.getItem("Number_of_comments");
  Number_of_liked_posts = sessionStorage.getItem("Number_of_liked_posts");
  Number_of_posts = sessionStorage.getItem("Number_of_posts");
};

export const post_component = () => {
  const img = new Image();
  img.src = `data:image/jpeg;base64,${sessionStorage.getItem("avatar")}`;
  return /*html*/ `<div class="lower-div">
  <main>
    <div id="c-post-modal" class="modal">
      <div class="modal-content">
          <div id="c-post-userinfo">
              <div class="c-post-pfp gender-${sessionStorage.getItem(
                "gender"
              )}"></div>
              <p id="c-post-nickname">${sessionStorage.getItem("username")}</p>
          </div>
          <textarea id="c-post-textArea"
              placeholder="What's on your mind?"></textarea>
          <div id="c-post-options">
              <div class="c-post-option">
                  <img src="${imgupload}" alt="upload Image"
                      title="upload Image" id="c-img-upload">
                  <input type="file" id="img-upload">
              </div>
              <div class="c-post-option">
                  <img src="${hashtag}" alt="Choose Category"
                      title="Choose Category" id="cat-choose-Btn">
              </div>
          </div>
          <div id="c-post-cats">
              <select id="c-post-cat-select">
                  <option class="c-option" value="1">General</option>
                  <option class="c-option" value="2">Engineering</option>
                  <option class="c-option" value="3">Travel</option>
                  <option class="c-option" value="4">Tech</option>
                  <option class="c-option" value="5">Mathematics</option>
              </select>
          </div>
          <div id="c-post-Btn">Create Post</div>
      </div>
    </div>
    <div id="contacts">
      <div id="profile">
        <div id="profile-name">Online Users</div>
      </div>
      <div id="c-title"></div>
      <div id="c-contacts"></div>
  </div>
    <div id="main_wrapper"></div>
  </main>

  <div class="side-divs">
    <div class="profile-card">
      <div class="profile-header">
        <div class="profileImage">
          <img src="${img.src}" style="width: 150px;
          height: 150px;
          border-radius: inherit;"alt="">
        </div>
      </div>
      <div class="UserInfo-div" id="UserInfo-div">
        <p class="UserName-p" style="font-size:20px">${sessionStorage.getItem(
          "username"
        )}</p>
        <div class="user-stats" style="font-size: 12px;">
        <p class="user-postd">Posts: ${Number_of_posts}</p>
        <p class="user-likes">Liked Posts: ${Number_of_liked_posts}</p>
        <p class="user-comments">Comments: ${Number_of_comments}</p>
        <p class="user-comments">Liked Comments: ${Number_of_liked_comments}</p>
      </div>
      </div>
    </div>
  </div>
</div>
`;
};

/**
 * returns a modal for all post information
 *
 * @returns {string} the html code for the post modal
 */
export const PostModal = async (postID) => {
  if (!sessionStorage.getItem("user_token")) {
    window.location.assign("/login");
    return;
  }

  let data = await fetch_post(postID);

  var gender = data.user.gender;
  let liked_img = noheart;
  if (data.liked) {
    liked_img = heart;
  }

  if (data) {
    document.getElementById("main_wrapper").insertAdjacentHTML(
      "afterbegin",
      /*html*/ `
  <div id="d-post-modal" class="modal">
  <div class="modal-content post-modal">
    <div id="d-post-wrapper">
      <div id="post-side">
        <div id="top-bar">
          <div id="author-profile">
            <div id="author-profile-img" class="author-${gender}">${data.user.username[0].toUpperCase()}</div>
            <div id="author-profile-info">
              <p id="author-name">${data.user.username}</p>
            </div>
          </div>
          <div id="post-creation-date">${new Date(
            data.creationDate
          ).toDateString()}</div>
        </div>
        <div id="post-content">
          ${data.content}
          ${
            data.Image_Path
              ? `<div class="p-image">
            <img src=${data.Image_Path} alt="post image">
          </div>`
              : ""
          }
        </div>
        <div class="p-stats">
          <div class="p-likeCount">
            <div class="p-likeBtn">
              <img src="${liked_img}" id="d-post-likeBtn" alt="like" />
            </div>
            <div class="p-likeStat" id="p-likeStat">${data.likes}</div>
          </div>
          <div class="p-commentCount">
            <img src="${comment}" alt="comment" />
            <div class="p-comment-Stat" id="p-comment-Stat-num">${
              data.number_of_comments
            }</div>
          </div>
        </div>
      </div>
      <div id="comment-side">
        <div id="comment-header">
          <div id="comment-header-text">Comments</div>
        </div>
        <div id="comments-wrapper"></div>
        <textarea id="comment-input" placeholder="Care to Comment?"></textarea>
        <button id="d-c-comment-btn">Create Comment!</button>
      </div>
    </div>
  </div>
  `
    );

    let post_modal = document.getElementById("d-post-modal");
    post_modal.style.display = "flex";

    // When user clicks outside window, remove modal. Account for c-post-modal too
    window.onclick = function (event) {
      if (event.target == post_modal) {
        post_modal.style.display = "none";
      } else if (event.target == document.getElementById("c-post-modal")) {
        document.getElementById("c-post-modal").style.display = "none";
      }
    };

    // handle post likes

    let likeBtn = document.getElementById("d-post-likeBtn");

    likeBtn.addEventListener("click", async () => {
      console.log("like-post");
      let post_id = postID;
      await RecordPostLikeEvent(post_id);
      await RedoStats();
      if (likeBtn.getAttribute("src") === noheart) {
        likeBtn.setAttribute("src", heart);

        document.getElementById("p-likeStat").innerText =
          parseInt(document.getElementById("p-likeStat").innerText) + 1;

        document
          .getElementById(`post_${postID}`)
          .querySelector(".p-likeBtn>img").src = heart;

        document
          .getElementById(`post_${postID}`)
          .getElementsByClassName("p-likeStat")
          .item(0).innerText = parseInt(
          document.getElementById("p-likeStat").innerText
        );
      } else {
        likeBtn.setAttribute("src", noheart);
        document.getElementById("p-likeStat").innerText =
          parseInt(document.getElementById("p-likeStat").innerText) - 1;

        document
          .getElementById(`post_${postID}`)
          .querySelector(".p-likeBtn>img").src = noheart;

        document
          .getElementById(`post_${postID}`)
          .getElementsByClassName("p-likeStat")
          .item(0).innerText = parseInt(
          document.getElementById("p-likeStat").innerText
        );
      }
    });

    document
      .getElementById("d-c-comment-btn")
      .addEventListener("click", async () => {
        console.log("submit-comment");
        let comment = document.getElementById("comment-input").value;
        let post_id = postID;
        await creat_comment({
          comment_text: comment,
          post_id: post_id,
        });

        document.getElementById("no-comments")
          ? document.getElementById("no-comments").remove()
          : "";

        document.getElementById("p-comment-Stat-num").innerText =
          parseInt(document.getElementById("p-comment-Stat-num").innerText) + 1;

        document
          .getElementById(`post_${postID}`)
          .getElementsByClassName("p-comment-Stat")
          .item(0).innerText = parseInt(
          document.getElementById("p-comment-Stat-num").innerText
        );

        document.getElementById("comment-input").value = "";
        render_comments(postID);

        await RedoStats();
      });

    render_comments(postID);
  }
};

export const Home = async () => {
  const myCookie = getCookie("session_id");
  if (!sessionStorage.getItem("user_token") && !myCookie) {
    window.location.assign("/login");
    return;
  }
  SetSessionStorageStats();
  await GetSessionStorageStats();
  document.getElementById("app").innerHTML = /*html*/ `
    ${LoadNav()}
    ${post_component()}
    `;

  // Modal Operations
  var modal = document.getElementById("c-post-modal");
  var modalOpenBtn = document.getElementById("c-post-start");

  if (modalOpenBtn && modal) {
    modalOpenBtn.onclick = function () {
      modal.style.display = "block";
    };
  }

  // When user clicks outside window, remove modal
  window.onclick = function (event) {
    if (event.target == modal) {
      modal.style.display = "none";
    }
  };

  const create_post_Btn = document.getElementById("c-post-Btn");

  if (create_post_Btn) {
    create_post_Btn.addEventListener("click", async () => {
      const post_text = document.getElementById("c-post-textArea").value;
      const raw_image_file = document.getElementById("c-img-upload").value;
      // stop removing the fix for category value capture
      let post_cat_arr = [];
      const post_category = parseInt(
        document.getElementById("c-post-cat-select").value
      );
      post_cat_arr.push(post_category);
      // end!!!
      const Image_Converstion_wrapper = async () => {
        return await convertImageToBase64(raw_image_file);
      };

      const postImage = await Image_Converstion_wrapper();

      const post_data = {
        user_token: sessionStorage.getItem("user_token"),
        post_text: post_text,
        post_image_base64: postImage,
        post_category: post_cat_arr,
      };

      try {
        const res = await fetch(BACKENDURL + "/post/create", {
          method: "POST",
          body: JSON.stringify(post_data),
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        });

        modal.style.display = "none";

        if (res.status === 201) {
          window.location.reload();
        } else {
          throw new Error(res.status, res.statusText);
        }
      } catch (error) {
        alert(error);
        console.error("post creation error", error);
      }
    });
  }
  let toggled = false;

  document.getElementById("cat-choose-Btn").addEventListener("click", () => {
    toggled = !toggled;
    if (toggled) {
      document.getElementById("c-post-cats").style.display = "block";
    } else {
      document.getElementById("c-post-cats").style.display = "none";
    }
  });

  document.getElementById("c-img-upload").addEventListener("click", () => {
    document.getElementById("img-upload").click();
  });

  const likeImages = document.querySelectorAll(".p-likeBtn img");

  likeImages.forEach((likeBtn) => {
    likeBtn.addEventListener("click", () => {
      if (likeBtn.getAttribute("src") === noheart) {
        likeBtn.setAttribute("src", heart);
        // add other like event
      } else {
        likeBtn.setAttribute("src", noheart);
        // add other unlike event
      }
    });
  });

  await OrgIndexPosts();
  ws.send(
    JSON.stringify({
      type: "get_dms",
    })
  );
};

export function getCookie(name) {
  // Create a string to search for the cookie name followed by an equal sign
  const nameEQ = name + "=";

  // Split the document.cookie string into an array of individual cookies
  const ca = document.cookie.split(";");

  // Loop through each cookie in the array
  for (let i = 0; i < ca.length; i++) {
    // Get the current cookie, trimming any leading whitespace
    let c = ca[i].trim();

    // Check if the current cookie starts with the name we are looking for
    if (c.indexOf(nameEQ) === 0) {
      // If so, return the value of the cookie (everything after the equal sign)
      return c.substring(nameEQ.length, c.length);
    }
  }

  // If the cookie was not found, return null
  return null;
}
