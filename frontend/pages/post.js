import { render_comments } from "../funcs/comments";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import comment from "../assets/comment.svg";
import { BACKENDURL } from "../funcs/vars";
import { LoadNav } from "../funcs/navbar";

const fetch_post = async (postId) => {
  const response = await fetch(`${BACKENDURL}/post/${postId}`, {
    credentials: "include",
  });
  const data = await response.json();
  return data;
};

export const Post = async () => {
  if (!sessionStorage.getItem("user_token")) {
    window.location.assign("/login");
    return;
  }
  let url = location.href;
  const urlParts = url.split("/");
  const postId = urlParts[urlParts.length - 1];

  let data = await fetch_post(postId);
  if (data) {
    await render_post_area();
    render_comments(postId);
    fill_in_post(data);

    const Like_Count_divs = document.querySelectorAll(".p-likeCount");
    Like_Count_divs.forEach((Like_Count_div) => {
      // Handle hearts
      const likeBtn = Like_Count_div.querySelector(".p-likeBtnP");
      likeBtn.addEventListener("click", handle_action_like);
      const likeCount = Like_Count_div.querySelector(".p-likeStat");
    });

    // Modal Operations
    var modal = document.querySelector(".modal");
    var modalcon = document.querySelector(".modal-content");
    var modalOpenBtn = document.querySelector(".p-commentCount");
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

    // Create comment button
    const create_com_Btn = document.getElementById("c-com-Btn");
    if (create_com_Btn) {
      create_com_Btn.addEventListener("click", async () => {
        const comment_text = document.getElementById("c-com-textArea").value;
        const comment_data = {
          user_token: sessionStorage.getItem("user_token"),
          post_id: postId,
          comment_text: comment_text,
        };
        await creat_comment(comment_data, modal);
      });
    }
  } else {
    console.error("No post data received");
    throw new Error(res.status, res.statusText);
  }
};

/**
 * Function to render the post area
 */

export const render_post_area = async () => {
  document.getElementById("app").innerHTML = /*html*/ `
${LoadNav()}
<main>
<div id="c-com-modal" class="modal">
<div class="modal-content">
  <!-- </div> -->
  <!-- <div class="modal-content"> -->
  <div id="c-com-userinfo">
    <div id="c-com-pfp">
      <img src="">
    </div>
    <p id="c-com-nickname"></p>
    <div class="p-creationDate"></div>
  </div>
  <div class="PostComment_Contaiar">
    <div class="com2ent">
      <div id="c-com-userinfo">
      </div>
      <textarea id="c-com-textArea" placeholder="What's on your mind?"></textarea>
    </div>
  </div>
  <div id="c-com-Btn">Add Comment</div>
</div>
</div>
<div id="post-page">
  <!-- for later (connectting the backend) -->
  <div id="post"></div>
  <div class="secDiv">
    <!-- Comments Section -->
    <h3 style="color:white;">Comments</h3>
    <div class="comments-section">
      <!-- for later (connectting the backend) -->
      <div id="comments"></div>
    </div>
  </div>
</div>
</main>
`;
};

/**
 * @param {*} data
 */

export const fill_in_post = async (data) => {
  const postDiv = document.getElementById("post");
  const content = document.querySelector(".p-content");
  const date = document.querySelector(".p-creationDate");
  let liked_img = noheart;
  let gender = data.user.gender;
  if (data.liked) {
    liked_img = heart;
  }
  console.log(data);
  if (data.image) {
    postDiv.innerHTML = /*html*/ `
                <div class="f-post" id="${data.id}">
                    <div class="p-header">
                        <div class="p-profileInfo">
                            <div class="p-profile-pic gender-${gender}"></div>
                            <div class="p-nickname">${data.user.username}</div>
                        </div>
                        <div class="p-creationDate">${new Date(
                          data.creationDate
                        ).toDateString()}</div>
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
                                <div class="p-likeBtnP">
                                    <img src="${liked_img}" id="${
      data.id
    }" alt="like"/>
                                </div>
                                <div class="p-likeStat" >${data.likes}</div>
                            </div>
                            <div class="p-commentCount">
                                <img src="${comment}" alt="comment" id="c-com-start" />
                                <div class="p-comment-Stat">${
                                  data.number_of_comments
                                }</div>
                            </div>
                        </div>
                        <div class="p-Category">
                        <p>#${data.category}</p>
                        </div>
                    </div>
                </div>
            `;
    content.innerHTML = `
                    <p style="margin: 0;">
                    ${data.content}
                    </p>
                    <div class="p-image">
                    </div>`;
    date.innerHTML = `${new Date(data.creationDate).toDateString()}`;
  } else {
    // postDiv = document.getElementById("post");
    // <!-- categories should be connected to the backend when it's done.  -->
    postDiv.innerHTML = /*html*/ `
              <div class="f-post noimage">
                  <div class="p-header">
                      <div class="p-profileInfo">
                          <div class="p-profile-pic gender-${gender}"></div>
                          <div class="p-nickname">${data.user.username}</div>
                      </div>
                      <div class="p-creationDate">${new Date(
                        data.creationDate
                      ).toDateString()}</div>
                  </div>
                  <div class="p-main">
                      <div class="p-content">
                          ${data.content}
                      </div>
                      <div class="p-stats">
                          <div class="p-likeCount">
                              <div class="p-likeBtnP">
                                  <img src="${liked_img}" id="${
      data.id
    }" alt="like"/>
                              </div>
                              <div class="p-likeStat">${data.likes}
                              </div>
                          </div>
                          <div class="p-commentCount">
                              <img src="${comment}" alt="comment" id="c-com-start" />
                              <div class="p-comment-Stat">${
                                data.number_of_comments
                              }
                              </div>
                          </div>
                          </div>
                          <div class="p-Category" style="padding:10px">
                          <p>#${data.category}</p>
                          </div>
                  </div>
              </div>
          `;
    content.innerHTML = `
          <p style="margin: 0;">
          </p>`;

    content.classList.add("noimage-c");
    date.innerHTML = `${new Date(data.creationDate).toDateString()}`;
  }
};

/**
 *
 * @param {*} comment_data
 * @param {*} modal
 */
export const creat_comment = async (comment_data, modal) => {
  const res = await fetch(BACKENDURL + "/comment/create", {
    method: "POST",
    body: JSON.stringify(comment_data),
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
};

export const handle_action_like = async (event) => {
  const postId = event.target.id;
  const likeBtn = event.target;
  // Get like button image

  const likeImgSrc = likeBtn.getAttribute("src");
  const like_counter_div = likeBtn.parentNode.nextElementSibling;
  // Toggle like button image
  await fetch_like_post_action_API(postId); // Toggle like or dislike API
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

export const fetch_like_post_action_API = async (postId) => {
  await fetch(`${BACKENDURL}/likePost/${postId}`, {
    method: "POST",
    credentials: "include",
  });
};
