import schat from "../assets/sendChat.svg";
import nchat from "../assets/nchat.svg";
import logout from "../assets/logout.svg";
import { NewChatMessage } from "../funcs/utils";
import { ws } from "../main";
import { Logout } from "../funcs/logout";

export const Chat = () => {
  document.getElementById("app").innerHTML += /*html*/ `
  <div id="contacts">
      <div id="profile">
        <div id="profile-name">${sessionStorage.getItem("username")}</div>
        <button type="button" title="New Chat" id="nchat">
          <img src="${logout}" alt="logout" title="logout" id="logout-btn"/>
        </button>
      </div>
      <div id="c-title"></div>
      <div id="c-contacts"></div>
  </div>
    <div id="messageArea">
      <div id="r-profile">
        <div id="pic"></div>
        <div id="r-name"></div>
      </div>
      <div id="message-space"></div>

      <div id="mbar">
        <textarea
          name="u-text"
          id="user-text"
          placeholder="Message..."
        ></textarea>
        <img src="${schat}" alt="DM" id="sendTextBtn" title="Send" />
      </div>
    </div>
  `;

  document
    .getElementById("logout-btn")
    .addEventListener("click", () => Logout());

  const messageInput = document.getElementById("user-text");

  document.getElementById("sendTextBtn").addEventListener("click", function () {
    const message = messageInput.value;
    if (!message.trim()) return;
    ws.send(
      JSON.stringify({
        type: "send_msg",
        req_Content: {
          sender: "",
          recipient: sessionStorage.getItem("chat_user"),
          msg_content: message,
        },
      })
    );
    NewChatMessage(message, true); // Add as self message
    messageInput.value = "";
  });
};
