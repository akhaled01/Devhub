import DM from "../images/nchat.svg"
import plus from "../images/plus.svg"

/**
 * 
 * @returns Navbar HTML
 */
export const LoadNav = () => {
  //FIXME - MAKE SURE THIS IS FIXED BEFORE PRODUCTION
  if (!localStorage.getItem("user_token")) {
    return /*html*/`
      <nav>
  <a href="/">
    <img class="navMainLogo" src="../images/logo.svg">
  </a>
  <ul class="actionitems">
    <li>
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
    return /*html*/`
      <nav>
        <a href="/">
          <img class="navMainLogo" src="../images/logo.svg">
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