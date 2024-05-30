import { AssemblePosts } from "./utils";
import { BACKENDURL } from "./vars";

/**
 * This function Sets Up the main page with
 * All the Posts By fetching all posts from
 * /post/all and invoking AssemblePosts
 * In order to assemble them in the main page
 */
export const OrgIndexPosts = async () => {
  const response = await fetch(BACKENDURL + "/post/all", {
    credentials: "include",
  });
  const posts_in_json = await response.json();
  AssemblePosts(posts_in_json);
};
