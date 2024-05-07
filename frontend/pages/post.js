import { LoadNav } from "../funcs/navbar";
import noheart from "../assets/unliked.svg";
import heart from "../assets/liked.svg";
import comment from "../assets/comment.svg";
import { BACKENDURL } from "../funcs/vars";

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
        })
        .catch((error) => {
            console.error("Error fetching post details:", error);
            // Handle error case
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

    async function fetchPost(postId) {
        if (postId === null) {
            console.error('postId is null');
            return;
        }
        const response = await fetch(`${BACKENDURL}/post/${postId}`, {
            credentials: "include",
        });
    
        const data = await response.json();

       const postDiv = document.getElementById("post");
        if (data.image) {
            postDiv.innerHTML = `
            <div class="f-post">
                <div class="p-header">
                    <div class="p-profileInfo">
                        <div class="p-profile-pic"></div>
                        <div class="p-nickname">${data.author}</div>
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
                            <img src="${comment}" alt="comment" />
                            <div class="p-comment-Stat">${data.number_of_comments}</div>
                        </div>
                    </div>
                </div>
            </div>
        `;
        } else {
            // postDiv = document.getElementById("post");
            // <!-- categories should be connected to the backend when it's done.  -->
            postDiv.innerHTML = `
          <div class="f-post noimage">
              <div class="p-header">
                  <div class="p-profileInfo">
                      <div class="p-profile-pic"></div>
                      <div class="p-nickname">${data.author}</div>
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
                          <img src="${comment}" alt="comment" />
                          <div class="p-comment-Stat">${data.number_of_comments}</div>
                      </div>
                  </div>
              </div>
          </div>
      `;
        }
    }
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