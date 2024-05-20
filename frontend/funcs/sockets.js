import { ws } from "../main";
import schat from "../assets/sendChat.svg";
import { NewChatMessage, sortByOnlineAndName } from "./utils";

export let CurrentChatUser = null;

/**
 *  Takes JSON data and assembles online users
 * @param {*} data
 */
export const AssembleOnlineUsersChat = (data) => {
  const contact_div = document.getElementById("c-contacts");
  if (!contact_div) return;
  contact_div.innerHTML = "";
  data.req_Content.forEach((user_obj) => {
    //let user = user_obj.user;
    console.log(user_obj);
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

      document
        .getElementById(user_obj.username)
        .addEventListener("click", () => {
          OrgChatHTML(user_obj.username);
        });
    }
  });
};

export const AssembleOnlineUsersIndex = (data) => {
  const list_div = document.getElementById("c-contacts");
  if (!list_div) return;
  list_div.innerHTML = "";
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

/**
 * Assembles the chat HTML in the index
 */
const OrgChatHTML = (username) => {
  const main_wrapper = document.getElementById("main_wrapper");
  main_wrapper.innerHTML = "";
  const recipient_div = document.createElement("div");
  recipient_div.id = "r-profile";
  recipient_div.innerHTML = username;
  main_wrapper.appendChild(recipient_div);

  const message_div = document.createElement("div");
  message_div.id = "message_space";
  main_wrapper.appendChild(message_div);

  const mdiv = document.createElement("div");
  mdiv.id = "mdiv";

  const message_input = document.createElement("textarea");
  message_input.id = "user-text";
  message_input.placeholder = "Write a msg :)";
  mdiv.appendChild(message_input);

  const send_btn = document.createElement("img");
  send_btn.src = schat;
  send_btn.id = "sendTextBtn";
  send_btn.title = "Send";
  mdiv.appendChild(send_btn);

  main_wrapper.appendChild(mdiv);

  document.getElementById("sendTextBtn").addEventListener("click", function () {
    const message = message_input.value;
    if (!message.trim()) return;
    ws.send(
      JSON.stringify({
        type: "send_msg",
        req_Content: {
          sender: "",
          recipient: username,
          msg_content: message,
        },
      })
    );
    NewChatMessage(message, true); // Add as self message
    message_input.value = "";
  });

  document.addEventListener('scroll', function(e) {
    if (e.target.scrollHeight - e.target.scrollTop <= e.target.clientHeight + 50) {
      ws.send(
        JSON.stringify({
          type: "get_msg",
          req_Content: {
            sender: "",
            recipient: username,
          },
        })
      );
    }
  });

  ws.send(
    JSON.stringify({
      type: "Open_chat",
      req_Content: {
        user_id: username,
      },
    })
  );
};
