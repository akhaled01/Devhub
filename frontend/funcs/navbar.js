import DM from "../assets/nchat.svg";
import plus from "../assets/plus.svg";

/**
 * Renders Navbar based on the if the user is logged in or
 * not
 *
 * @returns Navbar HTML
 */
export const LoadNav = () => {
  if (sessionStorage.getItem("user_token")) {
    if (!sessionStorage.getItem("user_token")) {
      window.location.assign("/login");
      return;
    }

    return /*html*/ `
      <nav>
  <a href="/">
    <img class="navMainLogo" src="../assets/logo.svg">
  </a>
  <ul class="actionitems">
    <li id="NewPost">
        <div class="actionItem" id="c-post-start">
          <img src="${plus}" alt="New Post" title="New Post">
        </div>
    </li>
    <li>
      <a href="/chat">
        <div class="actionItem">
          <img src="${DM}" alt="New Chat" title="Chat">
        </div>
      </a>
    </li>
  </ul>
  <div>
    <a href="/logout">
    <button id="btn-message" class="button-message">
     <!-- <div class="content-avatar gender-${sessionStorage.getItem("gender")}" id="c-avatar">
        <div class="status-user"></div>
      </div> -->
      <div class="notice-content">
        <div class="Logout-div">Logout</div>
        <div class="User-div">${sessionStorage.getItem("username")}</div>
      </div>
    </button>
    </a>
  </div>
</nav>
    `;

    // <button class="profile" id="profileBtn">Logout</button>
  } else {
    return /*html*/ `
      <nav>
        <a href="/">
          <img class="navMainLogo" src="../assets/logo.svg">
        </a>
        <div>
          <a href="/login">
            <button class="profile" id="profileBtn">Login</button>
          </a>
        </div>
      </nav>
    `;
  }
};
