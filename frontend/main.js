import { Logout } from "./funcs/logout";
import { ForumRouter } from "./funcs/router";

window.addEventListener("popstate", ForumRouter);

document.addEventListener("DOMContentLoaded", async () => {
  ForumRouter();
});
