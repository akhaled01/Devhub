import { ForumRouter } from "./funcs/router";
import { AssembleOnlineUsers } from "./funcs/sockets";
import { NewChatMessage } from "./funcs/utils";

export const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = () => {
  console.log("websocket Opening Successful");
};

ws.onclose = () => {
  console.log("websocket closing Successful");
};

// handle websocket events from backend
ws.onmessage = (e) => {
  console.log("RECIEVED MESSAGE:", JSON.parse(e.data));
  let data = JSON.parse(e.data);
  if (data.type === "message") {
    NewChatMessage(
      data.req_Content.msg_content,
      false,
      data.req_Content.sender
    );
  } else if (data.type === "online_user_list") {
    AssembleOnlineUsers(data);
  } else if (data.type === "open_chat_response") {
    // window.location.reload();
    let data = JSON.parse(e.data);
    document.getElementById("r-name").innerText =
      sessionStorage.getItem("chat_user");
    console.log(data.type);
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
