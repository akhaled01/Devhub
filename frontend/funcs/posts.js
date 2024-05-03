import { AssemblePosts } from "./utils";
import { BACKENDURL } from "./vars";

/**
 * This function Sets Up the main page with
 * All the Posts By fetching all posts from
 * /post/all and invoking AssemblePosts
 * In order to assemble them in the main page
 */

// Assuming you have the cookie value stored in a variable
const cookieValue = localStorage.getItem('user_token');

const requestOptions = {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    //'Cookie': `session_id=${cookieValue}`,
  },
  credentials: 'include', // This tells the browser to include cookies in cross-origin requests
  withCredentials: true,
};

export const OrgIndexPosts = async () => {
  const response = await fetch(BACKENDURL + "/post/all", requestOptions);
  const posts_in_json = await response.json();
  console.log(posts_in_json);
  AssemblePosts(posts_in_json);
};
