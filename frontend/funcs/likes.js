import { BACKENDURL } from "./vars";

export const RecordPostLikeEvent = async (post_id) => {
  await fetch(BACKENDURL + `/likePost/${post_id}`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
};
