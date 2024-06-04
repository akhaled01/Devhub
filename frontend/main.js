import { ForumRouter } from "./funcs/router";
import {
  AssembleOnlineUsersChat,
  AssembleOnlineUsersIndex,
  currentScrollHeight,
} from "./funcs/sockets";
import {
  NewChatMessage,
  PaginateHistoricalMessage,
} from "./funcs/utils";

// handle websocket connection
export const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = () => {
  console.log("websocket Opening Successful");
};

ws.onclose = () => {
  console.log("websocket closing Successful");
};

// handle websocket events from backend
ws.onmessage = (e) => {
  const data = JSON.parse(e.data);
  if (data.type === "message") {
    if (
      sessionStorage.getItem("chat_partner") &&
      data.req_Content.sender === sessionStorage.getItem("chat_partner")
    ) {
      NewChatMessage(
        data.req_Content.msg_content,
        data.req_Content.sender === sessionStorage.getItem("username"),
        data.req_Content.sender,
        new Date(data.req_Content.timestamp)
      );
    }
  } else if (data.type === "DMs") {
    if (document.getElementById("c-contacts")) {
      AssembleOnlineUsersChat(data);
    }
  } else if (data.type === "open_chat_response") {
    if (data.req_Content.length === 0) {
      sessionStorage.setItem("begin_id", 0);
      return;
    } else {
      sessionStorage.setItem("begin_id", data.req_Content[0].id);
    }
    data.req_Content.forEach((m) => {
      NewChatMessage(
        m.msg_content,
        m.sender === sessionStorage.getItem("username"),
        m.sender,
        new Date(m.timestamp)
      );
    });
  } else if (data.type === "typing_in_progress") {
    console.log(data);
    if (document.getElementById("r-profile")) {
      if (
        data.is_typing &&
        !document.getElementById("r-profile").querySelector(".typing-indicator") &&
        data.sender === sessionStorage.getItem("chat_partner")
      ) {
        document.getElementById("r-profile").innerHTML += `
      <div class="typing-indicator">
        <div class="typing-circle"></div>
        <div class="typing-circle"></div>
        <div class="typing-circle"></div>
      </div>
      `;
      } else {
        if (
          document
            .getElementById("r-profile")
            .querySelector(".typing-indicator")
        ) {
          document
            .getElementById("r-profile")
            .querySelector(".typing-indicator")
            .remove();
        }
      }
    }
  } else {
    let data = JSON.parse(e.data);
    if (data.req_Content.length === 0) {
      sessionStorage.setItem("begin_id", 0);
      return;
    } else {
      sessionStorage.setItem("begin_id", data.req_Content[0].id);
    }
    data.req_Content.forEach((m) => {
      PaginateHistoricalMessage(
        m.msg_content,
        m.sender === sessionStorage.getItem("username"),
        m.sender,
        new Date(m.timestamp)
      );
    });

    const messageBox = document.getElementById("message_space");

    const newScrollHeight = messageBox.scrollHeight;
    const scrollOffset = newScrollHeight - currentScrollHeight;

    // Restore the scroll position
    messageBox.scrollTop = scrollOffset;
  }
};

window.addEventListener("popstate", ForumRouter);
window.onbeforeunload = sessionStorage.removeItem("chat_user_selected");

document.addEventListener("DOMContentLoaded", async () => {
  ForumRouter();
});
