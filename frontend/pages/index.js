import { LoadNav } from "../funcs/navbar";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import comment from "../assets/comment.svg";
import imgupload from "../assets/imageupload.svg";
import hashtag from "../assets/hashtag.svg";
import { OrgIndexPosts } from "../funcs/posts";
import { BACKENDURL } from "../funcs/vars";

export const Home = async () => {
  if (!localStorage.getItem("user_token")) {
    window.location.assign("/login")
    return
  }

  let username = localStorage.getItem("username");
  let avatar = localStorage.getItem("avatar");

  document.getElementById("app").innerHTML = /*html*/ `
    ${LoadNav()}
    <main>
    <div id="c-post-modal" class="modal">
        <div class="modal-content">
            <div id="c-post-userinfo">
                <div id="c-post-pfp">
                    <img src="${avatar}">
                </div>
                <p id="c-post-nickname">${username}</p>
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

  // Check if the elements exist before attaching event handlers
  if (modalOpenBtn && modal) {
    // When the user clicks the button, open the modal
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


  // document.addEventListener('DOMContentLoaded', () => {
  // Get the "NewPost" div element
  const newPostDiv = document.getElementById('c-post-Btn');

  // Check if the element exists before adding the event listener
  if (newPostDiv) {
    // Add a click event listener to the "NewPost" div
    newPostDiv.addEventListener('click', () => {

      // Get the post text and image from the form
      const postText = document.getElementById('c-post-textArea').value;
      const postImage = ''; // Get the base64-encoded image data
      const postCategory = document.getElementById('cat-choose-Btn').value;

      // Create an object with the post data
      const postData = {
        post_text: postText,
        post_image_base64: postImage,
        post_category: postCategory, // Set the desired category ID
        creator_id: localStorage.getItem("user_id"), // Get the user ID from local storage
      };

      // Send a POST request to the backend
      fetch(BACKENDURL + '/post/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: "include",

        body: JSON.stringify(postData),
      })
        .then(response => {
          modal.style.display = "none";
          if (response.ok) {
            // Post created successfully
            console.log('Post created successfully');
            // Redirect to the post page or update the UI as needed
          } else {
            // Handle error response
            console.error('Error creating post');
          }
        })
        .catch(error => {
          console.error('Error creating post:', error);
        });
    });
  }
  // });


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

  // liking event listener
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

  await OrgIndexPosts();
};

export async function fetchPost() {
  const response = await fetch(BACKENDURL + `/post/all`);
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
                      ${data.Image_Path
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
