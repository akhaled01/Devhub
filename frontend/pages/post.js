import { fetchComments } from "../funcs/comments";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import comment from "../assets/comment.svg";
import { BACKENDURL } from "../funcs/vars";
import { LoadNav } from "../funcs/navbar";

/**
 * This function fetches the main post page
 * (the one with all the details like comments lol)
 */


export const Post = () => {
    if (!sessionStorage.getItem("user_token")) {
        window.location.assign("/login");
        return;
    }
    let url = location.href;
    const urlParts = url.split("/");
    // const urlParams = new URLSearchParams(window.location.search);
    const postId = urlParts[urlParts.length - 1];
    console.log(urlParts, "------", postId);

    console.log(`Fetching post details for post ID: ${postId}`);

    // if (!postId) {
    //     // Handle case when post ID is not provided
    //     window.location.assign("/");
    //     return;
    // }

    // Fetch post details based on the postId
    fetch(`${BACKENDURL}/post/${postId}`, {
        credentials: "include",
    }
        // .then((response) => response.json())
    )

        .then((postData) => {
            document.getElementById("app").innerHTML = /*html*/ `
    ${LoadNav()}
    <main>
    <div id="c-com-modal" class="modal">
    <div class="modal-content">
      <!-- </div> -->
      <!-- <div class="modal-content"> -->
      <div id="c-com-userinfo">
        <div id="c-com-pfp">
          <img src="${sessionStorage.getItem("avatar")}">
        </div>
        <p id="c-com-nickname">${sessionStorage.getItem(
                "username"
            )}</p>
        <div class="p-creationDate"></div>
      </div>
      <div class="p-content">
      </div>
      <div class="PostComment_Contaiar">
        <div class="shakta"></div>
        <div class="com2ent">
          <div id="c-com-userinfo">
            <div id="c-com-pfp">            
              <img src="${sessionStorage.getItem("avatar")}" alt="">
            </div>
            <p id="c-com-nickname">${sessionStorage.getItem(
                "username"
            )}</p>
          </div>
          <textarea id="c-com-textArea" placeholder="What's on your mind?"></textarea>
        </div>
      </div>
      <div id="c-com-Btn">Create a Replay</div>
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
            fetchPost(postId);
            fetchComments(postId);
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

            async function fetchPost(postId) {
                if (postId === null) {
                    console.error('postId is null');
                    return;
                }
                const response = await fetch(`${BACKENDURL}/post/${postId}`, {
                    credentials: "include",
                });

                const data = await response.json();

                console.log(data);
                // load the div that we need to put thr data in
                const postDiv = document.getElementById("post");
                const content = document.querySelector(".p-content");
                const date = document.querySelector(".p-creationDate");
                if (data.image) {
                    postDiv.innerHTML = `
            <div class="f-post">
                <div class="p-header">
                    <div class="p-profileInfo">
                        <div class="p-profile-pic"></div>
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
                            <div class="p-likeBtn">
                                <img src="${noheart}" alt="like"/>
                            </div>
                            <div class="p-likeStat">${data.likes}</div>
                        </div>
                        <div class="p-commentCount">
                            <img src="${comment}" alt="comment" id="c-com-start" />
                            <div class="p-comment-Stat">${data.number_of_comments}</div>
                        </div>
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
                    date.innerHTML = `${new Date(
                        data.creationDate
                    ).toDateString()}`;
                } else {
                    // postDiv = document.getElementById("post");
                    // <!-- categories should be connected to the backend when it's done.  -->
                    postDiv.innerHTML = `
          <div class="f-post noimage">
              <div class="p-header">
                  <div class="p-profileInfo">
                      <div class="p-profile-pic"></div>
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
                          <div class="p-likeBtn">
                              <img src="${noheart}" alt="like" />
                          </div>
                          <div class="p-likeStat">${data.likes}</div>
                      </div>
                      <div class="p-commentCount">
                          <img src="${comment}" alt="comment" id="c-com-start" />
                          <div class="p-comment-Stat">${data.number_of_comments}</div>
                      </div>
                  </div>
              </div>
          </div>
      `;
                    content.innerHTML = `        
      <p style="margin: 0;">
      ${data.content}
      </p>`;

      content.classList.add("noimage-c")
                    date.innerHTML = `${new Date(
                        data.creationDate
                    ).toDateString()}`;

                }
                // Modal Operations
                // var modal = document.getElementById("c-com-modal");
                var modal = document.querySelector(".modal");
                var modalcon = document.querySelector(".modal-content");
                if (!data.image && modalcon) {
                    modalcon.style.minHeight = "260px";
                }
                console.log(modal);
                var modalOpenBtn = document.querySelector(".p-commentCount");
                // var modali = document.querySelector(".f-post");
                // var modalOpenBtn = document.getElementById("c-com-start");
                console.log(modalOpenBtn, modal);
                if (modalOpenBtn && modal) {
                    console.log("-------------");
                    modalOpenBtn.onclick = function () {
                        modal.style.display = "block";
                        console.log("Modal opened");
                    };
                }

                // When user clicks outside window, remove modal
                window.onclick = function (event) {
                    if (event.target == modal) {
                        modal.style.display = "none";
                    }
                };

                const create_com_Btn = document.getElementById("c-com-Btn");

                if (create_com_Btn) {
                    create_com_Btn.addEventListener("click", async () => {
                        const comment_text = document.getElementById("c-com-textArea").value;

                        const comment_data = {
                            user_token: sessionStorage.getItem("user_token"),
                            post_id: postId,
                            comment_text: comment_text
                        };

                        try {
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
                        } catch (error) {
                            alert(error);
                            console.error("post creation error", error);
                        }
                    });
                }
            }


        })
        .catch((error) => {
            console.error("Error fetching post details:", error);
            // Handle error case
        });
};


//     <div class="f-post">
//     <div class="p-header">
//         <div class="p-profileInfo">
//             <div class="p-profile-pic"></div>
//             <div class="p-nickname">Ralph</div>
//         </div>
//         <div class="p-creationDate">2 Hours Ago</div>
//     </div>
//     <div class="p-main">
//         <div class="p-content">
//             Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptas id
//             unde quisquam enim ullam ex quaerat velit numquam autem temporibus.
//             Aut ex vel necessitatibus, optio maxime debitis! Quo, inventore
//             ducimus!
//             <div class="p-image">
//                 <img src="https://images.pexels.com/photos/8434281/pexels-photo-8434281.jpeg?auto=compress&amp;cs=tinysrgb&amp;w=1260&amp;h=750&amp;dpr=2"
//                     alt="stuff">
//             </div>
//         </div>
//         <div class="p-stats">
//             <div class="p-likeCount">
//                 <div class="p-likeBtn">
//                     <img src="${noheart}" alt="like" />
//                 </div>
//                 <div class="p-likeStat">9</div>
//             </div>
//             <div class="p-commentCount">
//                 <img src="${comment}" alt="comment" />
//                 <div class="p-comment-Stat">10</div>
//             </div>
//         </div>
//     </div>
// </div>
//     <div class="comment">
//     <div class="comment-header">
//       <div class="c-profileInfo">
//         <div class="c-profile-pic">
//           <!-- <img src="user1-avatar.png" alt="User 1 Avatar" class="user-avatar"> -->
//         </div>
//         <div class="c-nickname">james_of_pdx</div>
//       </div>
//       <div class="c-creationDate">2 Hours Ago</div>
//     </div>
//     <div class="p-main">Nice but why?
//       <div class="p-stats">
//         <div class="p-likeCount">
//           <div class="p-likeBtn">
//             <img src="${noheart}" alt="like" />
//           </div>
//           <div class="p-likeStat">9</div>
//         </div>
//       </div>
//     </div>
//   </div>
//   <div class="comment">
//     <div class="comment-header">
//       <div class="c-profileInfo">
//         <div class="c-profile-pic">
//           <!-- <img src="user1-avatar.png" alt="User 1 Avatar" class="user-avatar"> -->
//         </div>
//         <div class="c-nickname">hatter1</div>
//       </div>
//       <div class="c-creationDate">3 Hours Ago</div>
//     </div>
//     <div class="p-main">Ew Coffe
//       <div class="p-stats">
//         <div class="p-likeCount">
//           <div class="p-likeBtn">
//             <img src="/images/unliked.svg" alt="like">
//           </div>
//           <div class="p-likeStat">9</div>
//         </div>
//       </div>
//     </div>
//   </div>
//   <div class="comment">
//     <div class="comment-header">
//       <div class="c-profileInfo">
//         <div class="c-profile-pic">
//           <!-- <img src="user1-avatar.png" alt="User 1 Avatar" class="user-avatar"> -->
//         </div>
//         <div class="c-nickname">hatter2</div>
//       </div>
//       <div class="c-creationDate">2 Hours Ago</div>
//     </div>
//     <div class="p-main">yea, EW Coffe
//       <div class="p-stats">
//         <div class="p-likeCount">
//           <div class="p-likeBtn">
//             <img src="/images/unliked.svg" alt="like">
//           </div>
//           <div class="p-likeStat">2</div>
//         </div>
//       </div>
//     </div>
//   </div>
// </div>

// <!-- post the comment. -->
// <div class="modal-contentt">
//   <div id="c-post-userinfo">
//     <div id="c-post-pfp"></div>
//     <p id="c-post-nickname" style="color: white;">_.ak79</p>
//   </div>
//   <textarea id="c-post-textArea" placeholder="What's on your mind?"></textarea>
//   <div id="c-post-Btn">Create Post</div>
// </div>
// </div>