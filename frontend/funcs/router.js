import { Home } from "../pages";
import { Login } from "../pages/auth/login";
import { Signup } from "../pages/auth/signup";
import { Chat } from "../pages/chat";
import { Post } from "../pages/post";
import { Logout } from "./logout";
import { UpdateCSS } from "./utils";

const routes = {
  "/post": { component: Post, stylesheet: "/css/post.css" },
  "/signup": { component: Signup, stylesheet: "/css/auth.css" },
  "/login": { component: Login, stylesheet: "/css/auth.css" },
  "/chat": { component: Chat, stylesheet: "/css/chat.css" },
  "/": { component: Home, stylesheet: "/css/index.css" },
};

const ExtractHref = () => {
  let url = location.href;
  const urlParts = url.split("/");
  if (urlParts.length > 4) {
    var pathname = urlParts[urlParts.length - 2];
  } else {
    var pathname = urlParts[urlParts.length - 1];
  }
  return pathname;
};

/**
 * frontend router
 */
export const ForumRouter = () => {
  const path = ExtractHref();
  const route = routes["/" + path];
  if (path === "logout") {
    Logout();
    return;
  }
  if (route) {
    route.component();
    UpdateCSS(route.stylesheet);
  } else {
    Home(); // Default route
    SetSessionStorageStats();
    UpdateCSS("/css/index.css"); // Default stylesheet
  }
};
