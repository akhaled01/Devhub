import { ws } from "../main";

export let CurrentChatUser = null;

/**
 * Registers a new user chat sess
 * @param {*} username
 */
export const SaveCurrentChatUser = async (username) => {
  console.log(username);
  if (
    sessionStorage.getItem("chat_user_selected") &&
    username === sessionStorage.getItem("chat_user")
  )
    return;
  sessionStorage.setItem("chat_user", username);
  sessionStorage.setItem("chat_user_selected", "true");
  document.getElementById("message-space").innerHTML = "";
  ws.send(
    JSON.stringify({
      type: "Open_chat",
      req_Content: {
        user_id: sessionStorage.getItem("chat_user"),
      },
    })
  );
  CurrentChatUser = username;
};

/**
 *  Takes JSON data and assembles online users
 * @param {*} data
 */
export const AssembleOnlineUsers = (data) => {
  const contact_div = document.getElementById("c-contacts");
  contact_div.innerHTML = "";
  console.log(data);
  data.req_Content.forEach((user_obj) => {
    let user = user_obj.user;
    console.log(user_obj.is_online);
    if (user.username !== sessionStorage.getItem("username")) {
      const contactDiv = document.createElement("div");
      contactDiv.classList.add("contact");
      contactDiv.id = user.username;

      const nameDiv = document.createElement("div");
      nameDiv.classList.add("name");
      nameDiv.textContent = user.username;

      contactDiv.appendChild(nameDiv);
      contact_div.appendChild(contactDiv);

      document.getElementById(user.username).addEventListener("click", () => {
        SaveCurrentChatUser(user.username);
      });
    }
  });
};
