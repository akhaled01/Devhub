export const RecordLikes = async (event = 1, post_id) => {
  const response = await fetch(`/api/likes?Post_id=${post_id}`, {
    method: "POST",
    body: JSON.stringify({
      Post_id: post_id,
      User_id: event,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });
  const data = await response.json();
  return data;
};
