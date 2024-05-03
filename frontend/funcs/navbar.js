import DM from "../assets/nchat.svg";
import plus from "../assets/plus.svg";

/**
 * Renders Navbar based on the if the user is logged in or
 * not
 *
 * @returns Navbar HTML
 */
export const LoadNav = () => {
  if (localStorage.getItem("user_token")) {
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
  <div class="filters">
    <label class="filter">
      <input type="radio" name="radio" checked="">
      <span class="name">Newest</span>
    </label>
    <label class="filter">
      <input type="radio" name="radio">
      <span class="name">Liked By You</span>
    </label>
  </div>
  <div>
    <a href="/logout">
      <button class="profile" id="profileBtn">Logout</button>
    </a>
  </div>
</nav>
    `;
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



