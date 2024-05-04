import schat from "../assets/sendChat.svg";
import nchat from "../assets/nchat.svg";
import { NewChatWS } from "../funcs/sockets";
import { NewChatMessage } from "../funcs/utils";

export const Chat = () => {
  document.getElementById("app").innerHTML += `
  <div id="contacts">
      <div id="profile">
        <div id="profile-name">${sessionStorage.getItem("username")}</div>
        <button type="button" title="New Chat" id="nchat">
          <img src="${nchat}" alt="new chat" />
        </button>
      </div>

      <div id="c-title"></div>

      <div class="contact">
        <div class="profile-pic">
          <div class="pfp"></div>
        </div>
        <div class="name">sahmed</div>
      </div>
    </div>
    <div id="messageArea">
      <div id="r-profile">
        <div id="pic"></div>
        <div id="r-name">sahmed</div>
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

  const messageInput = document.getElementById("user-text");
  let ws = NewChatWS();

  document.getElementById("sendTextBtn").addEventListener("click", function () {
    const message = messageInput.value;
    if (!message.trim()) return;
    ws.send(
      JSON.stringify({
        type: "send_msg",
        req_Content: {
          sender: "",
          recipient: "VK", //TODO: Make it change based on session
          msg_content: message,
        },
      })
    );
    NewChatMessage(message, true); // Add as self message
    messageInput.value = "";
  });
};
