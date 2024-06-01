import { ForumRouter } from "./funcs/router";
import {
  AssembleOnlineUsersChat,
  AssembleOnlineUsersIndex,
} from "./funcs/sockets";
import { NewChatMessage } from "./funcs/utils";

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
    NewChatMessage(
      data.req_Content.msg_content,
      data.req_Content.sender === sessionStorage.getItem("username"),
      data.req_Content.sender,
      new Date(data.req_Content.timestamp)
    );
  } else if (data.type === "DMs") {
    if (document.getElementById("c-contacts")) {
      AssembleOnlineUsersChat(data);
    } else {
      AssembleOnlineUsersIndex(data);
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
  } else {
    let data = JSON.parse(e.data);
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
  }
};

window.addEventListener("popstate", ForumRouter);
window.onbeforeunload = sessionStorage.removeItem("chat_user_selected");

document.addEventListener("DOMContentLoaded", async () => {
  ForumRouter();
});
