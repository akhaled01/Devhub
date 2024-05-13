import { ws } from "../main";
import { sortByOnlineAndName } from "./utils";

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
export const AssembleOnlineUsersChat = (data) => {
  const contact_div = document.getElementById("c-contacts");
  // let mdata = sortByOnlineAndName(data);
  if (!contact_div) return;
  contact_div.innerHTML = "";
  console.log(data);
  data.req_Content.forEach((user_obj) => {
    //let user = user_obj.user;
    if (user_obj.username !== sessionStorage.getItem("username")) {
      const contactDiv = document.createElement("div");
      contactDiv.classList.add("contact");
      contactDiv.id = user_obj.username;

      const nameDiv = document.createElement("div");
      nameDiv.classList.add("name");
      nameDiv.textContent = user_obj.username;

      contactDiv.appendChild(nameDiv);
      contact_div.appendChild(contactDiv);

      if (user_obj.is_online) {
        nameDiv.classList.add("online");
      }

      document.getElementById(user_obj.username).addEventListener("click", () => {
        SaveCurrentChatUser(user_obj.username);
      });
    }
  });
};

export const AssembleOnlineUsersIndex = (data) => {
  const list_div = document.getElementById("online-user-list");
  //let mdata = sortByOnlineAndName(data);
  console.log("MDATA",data);
  if (!list_div) return;
  list_div.innerHTML = "";
  console.log(data);
  data.req_Content.forEach((user_obj) => {
    let user = user_obj;
    if (user.username !== sessionStorage.getItem("username")) {
      const user_div = document.createElement("li");
      user_div.textContent = user.username;
      user_div.id = user.username;
      if (user_obj.is_online) {
        user_div.classList.add("online");
      }
      list_div.appendChild(user_div);

      document.getElementById(user.username).addEventListener("click", () => {
        SaveCurrentChatUser(user.username);
        window.location.assign("/chat");
      });
    }
  });
};
