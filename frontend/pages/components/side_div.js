// Side Div is the div that contains the profile card and the online user section.
export const side_div = (img_src, username, Number_of_posts, Number_of_liked_posts, Number_of_comments, Number_of_liked_comments) => {
    return /*html*/ `
    ${profile_card(
        img_src,
        username,
        Number_of_posts,
        Number_of_liked_posts,
        Number_of_comments,
        Number_of_liked_comments)}
    ${online_user_section()}
  `
}

// Online User Section is the div that contains the online users.
export const online_user_section = () => {
    return /*html*/ `
    <div class="online-user-section">
      <h2 class="online-text">Users</h2>
      <ul class="user-list" id="online-user-list"> </ul>
    </div>
    `
}

// Profile Card is the div that contains the profile card and user stats
export const profile_card = (img_src, username, Number_of_posts, Number_of_liked_posts, Number_of_comments, Number_of_liked_comments) => {
    return /*html*/ `
    <div class="profile-card">
      <div class="profile-header">
        <div class="profileImage">
          <img src="${img_src}" style="width: 150px;
          height: 150px;
          border-radius: inherit;"alt="">
        </div>
      </div>
      <div class="UserInfo-div">
        <p class="UserName-p" style="font-size:20px">${username}</p>
        <div class="user-stats" style="font-size: 12px;">
        <p class="user-postd">Posts: ${Number_of_posts}</p>
        <p class="user-likes">Liked Posts: ${Number_of_liked_posts}</p>
        <p class="user-comments">Comments: ${Number_of_comments}</p>
        <p class="user-comments">Liked Comments: ${Number_of_liked_comments}</p>
        </div>
      </div>
    </div>
    `
}