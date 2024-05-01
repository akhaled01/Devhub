import { LoadNav } from "../funcs/navbar";
import noheart from "../images/unliked.svg";
import heart from "../images/liked.svg";
import comment from "../images/comment.svg";
import imgupload from "../images/imageupload.svg";
import hashtag from "../images/hashtag.svg";

export const Home = () => {
  document.getElementById("app").innerHTML = /*html*/ `
    ${LoadNav()}
    <main>
    <div id="c-post-modal" class="modal">
        <div class="modal-content">
            <div id="c-post-userinfo">
                <div id="c-post-pfp"></div>
                <p id="c-post-nickname">_.ak79</p>
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
                    <option class="c-option" value="general">General</option>
                    <option class="c-option" value="general">Dev</option>
                    <option class="c-option" value="general">News</option>
                    <option class="c-option" value="general">Non
                        Fiction</option>
                </select>
            </div>
            <div id="c-post-Btn">Create Post</div>
        </div>
    </div>
    <div id="posts"></div>
</main>
  `;

  // Modal Operations
  var modal = document.getElementById("c-post-modal");
  var modalOpenBtn = document.getElementById("c-post-start");

  // When the user clicks the button, open the modal
  modalOpenBtn.onclick = function () {
    modal.style.display = "block";
  };

  // When user clicks outside window, remove modal
  window.onclick = function (event) {
    if (event.target == modal) {
      modal.style.display = "none";
    }
  };

  // add event listener for category button
  let toggled = false;

  document.getElementById("cat-choose-Btn").addEventListener("click", () => {
    toggled = !toggled;
    if (toggled) {
      document.getElementById("c-post-cats").style.display = "block";
    } else {
      document.getElementById("c-post-cats").style.display = "none";
    }
  });

  // add event listener to hidden file upload

  document.getElementById("c-img-upload").addEventListener("click", () => {
    document.getElementById("img-upload").click();
  });

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
    });
  });
};

export async function fetchPost() {
  const response = await fetch(`/post/all`);
  const data = await response.json();

  const postDiv = document.getElementById("post");
  if (data.image) {
    postDiv.innerHTML = `
        <div class="f-post" ${!data.Image_Path ? "noimage" : ""}>
              <div class="p-header">
                  <div class="p-profileInfo">
                      <div class="p-profile-pic"></div>
                      <div class="p-nickname">${data.author}</div>
                  </div>
                  <div class="p-creationDate">${data.creationDate}</div>
              </div>
              <div class="p-main">
                  <div class="p-content">
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
                              <img src="${noheart}" alt="like" />
                          </div>
                          <div class="p-likeStat">${data.likes}</div>
                      </div>
                      <div class="p-commentCount">
                          <img src="${comment}" alt="comment" />
                          <div class="p-comment-Stat">${data.comments}</div>
                      </div>
                  </div>
              </div>
          </div>
      `;
  }
}
// categories for later.
/* export async function fetchPost() {
  const response = await fetch(``);
  const data = await response.json();
  const postDiv = document.getElementById("post");
} */