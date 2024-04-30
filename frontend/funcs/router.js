import { Home } from "../pages";
import { Login } from "../pages/auth/login";
import { Signup } from "../pages/auth/signup"
import { Chat } from "../pages/chat";
import { Post } from "../pages/post";
import { UpdateCSS } from "./utils";

const routes = {
  '/post': { component: Post, stylesheet: '/css/post.css' },
  '/signup': { component: Signup, stylesheet: '/css/auth.css' },
  '/login': { component: Login, stylesheet: '/css/auth.css' },
  '/chat': { component: Chat, stylesheet: '/css/chat.css' },
  '/': { component: Home, stylesheet: '/css/index.css' },
};

const ExtractHref = () => {
  let url = location.href
  const urlParts = url.split("/");

  const pathname = urlParts[urlParts.length - 1];

  return pathname
}

/**
 * frontend router
 */
export const ForumRouter = () => {
  const path = ExtractHref();
  const route = routes["/" + path];
  if (route) {
    route.component();
    UpdateCSS(route.stylesheet);
  } else {
    Home(); // Default route
    UpdateCSS('/css/index.css'); // Default stylesheet
  }
};