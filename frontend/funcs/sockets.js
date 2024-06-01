import { ws } from "../main";
import schat from "../assets/sendChat.svg";
import { NewChatMessage, sortByOnlineAndName } from "./utils";
import { GetSessionStorageStats } from "../pages";

export let CurrentChatUser = null;

/**
 *  Takes JSON data and assembles online users
 * @param {*} data
 */
export const AssembleOnlineUsersChat = (data) => {
  const contact_div = document.getElementById("c-contacts");
  if (!contact_div) return;
  contact_div.innerHTML = "";
  console.log(data);
  data.req_Content.forEach((user_obj) => {
    if (user_obj.username !== sessionStorage.getItem("username")) {
      const contactDiv = document.createElement("div");
      contactDiv.classList.add("contact");
      contactDiv.id = user_obj.username;

      const nameDiv = document.createElement("div");
      nameDiv.classList.add("name");
      nameDiv.textContent = user_obj.username;

      if (user_obj.is_online) {
        nameDiv.classList.add("online");
      } else {
        nameDiv.classList.add("offline");
      }

      if (
        user_obj.msg_status === false &&
        user_obj.last_message_time !== "" &&
        user_obj.msg_sender !== sessionStorage.getItem("username")
      ) {
        const redCircle = document.createElement("div");
        redCircle.classList.add("red-circle");
        nameDiv.appendChild(redCircle);
      }

      contactDiv.appendChild(nameDiv);
      contact_div.appendChild(contactDiv);

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
    if (user_obj.username !== sessionStorage.getItem("username")) {
      const user_div = document.createElement("li");
      user_div.textContent = user_obj.username;
      user_div.id = user_obj.username;

      if (user_obj.is_online) {
        user_div.classList.add("online");
      } else {
        user_div.classList.add("offline");
      }

      list_div.appendChild(user_div);

      document
        .getElementById(user_obj.username)
        .addEventListener("click", () => {
          SaveCurrentChatUser(user_obj.username);
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

  message_input.addEventListener("input", () => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(
        JSON.stringify({
          type: "typing-event",
          req_Content: {
            user_id: username,
          },
        })
      );
    } else {
      console.error("WebSocket is not open. Ready state is:", ws.readyState);
    }
  });
  
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

  const message_Box = document.getElementById("message_space");
  message_Box.addEventListener("scroll", function (e) {
    if (e.target.scrollTop === 0) {
      let beginid = parseInt(sessionStorage.getItem("begin_id"));
      ws.send(
        JSON.stringify({
          type: "load_messages",
          req_Content: {
            User_id: username,
            Begin_id: beginid,
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
