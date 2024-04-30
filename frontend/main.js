import { ForumRouter } from "./funcs/router";

window.addEventListener('popstate', ForumRouter);

document.addEventListener('DOMContentLoaded', () => {
  ForumRouter();
});