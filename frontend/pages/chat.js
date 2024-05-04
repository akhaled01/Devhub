import schat from "../assets/sendChat.svg";
import nchat from "../assets/nchat.svg";

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
      <div class="contact">
        <div class="profile-pic">
          <div class="pfp"></div>
        </div>
        <div class="name">malsamma</div>
      </div>
    </div>
    <div id="messageArea">
      <div id="r-profile">
        <div id="pic"></div>
        <div id="r-name">sahmed</div>
      </div>
      <div id="message-space">
        <div class="m">
          <div class="message">
            <div class="sender-info">
              <div class="sname">Sahmed</div>
              <div class="date">2 Hours Ago</div>
            </div>
            <p>hi</p>
          </div>
        </div>

        <div class="mself">
          <div class="message self">
            <div class="sender-info">
              <div class="sname">You</div>
              <div class="date">2 Hours Ago</div>
            </div>
            <p>hi</p>
          </div>
        </div>

        <div class="m">
          <div class="message">
            <div class="sender-info">
              <div class="sname">Sahmed</div>
              <div class="date">2 Hours Ago</div>
            </div>
            <p>hi</p>
          </div>
        </div>

        <div class="mself">
          <div class="message self">
            <div class="sender-info">
              <div class="sname">You</div>
              <div class="date">2 Hours Ago</div>
            </div>
            <p>
              Lorem ipsum dolor, sit amet consectetur adipisicing elit. Ab,
              fugit. Corrupti accusantium eligendi magnam blanditiis, obcaecati
              consequatur perferendis quaerat qui dolor aperiam repudiandae
              nobis, sit laborum dignissimos deserunt voluptate facilis?
            </p>
          </div>
        </div>
      </div>

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

  const chatArea = document.getElementById("message-space");
  const messageInput = document.getElementById("user-text");
  let ws = new WebSocket("ws://localhost:8080/ws");

  document.addEventListener("DOMContentLoaded", () => {
    ws.onopen = () => {
      console.log("websocket Opening Successful");
    };

    ws.onclose = () => {
      console.log("websocket closing Successful");
    };

    ws.onmessage = (e) => {
      alert("message came");
      addMessage(e.req_Content.msg_content, false, e.req_Content.sender);
    };
  });

  /**
   *
   * Function to add a message div via websocket
   *
   * @param {*} message - text to be sent
   * @param {*} isSelf - mark true if the current user is sending
   * @param {string} [name=""] - name of sender (only if isSelf is false)
   */
  function addMessage(message, isSelf, name = "") {
    const messageElement = document.createElement("div");
    const actualMessage = document.createElement("div");
    // is self checks if the message came from the current user, not the
    // other one
    if (isSelf) {
      messageElement.classList.add("mself");
      actualMessage.classList.add("self");
      actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">You</div>
              <div class="date">2 Hours Ago</div>
            </div>`;
    } else {
      messageElement.classList.add("m");
      actualMessage.innerHTML += `<div class="sender-info">
              <div class="sname">${name}</div>
              <div class="date">2 Hours Ago</div>
            </div>`;
    }

    actualMessage.classList.add("message");

    const content = document.createElement("p");
    content.textContent = message;
    actualMessage.appendChild(content);
    messageElement.appendChild(actualMessage);
    chatArea.appendChild(messageElement);
    chatArea.scrollTop = chatArea.scrollHeight; // Scroll to bottom
  }

  document.getElementById("sendTextBtn").addEventListener("click", function () {
    const message = messageInput.value;
    if (!message.trim()) return;
    ws.send(
      JSON.stringify({
        type: "send_msg",
        req_Content: {
          sender: "",
          recipient: "VK",
          msg_content: message,
        },
      })
    );
    addMessage(message, true); // Add as self message
    messageInput.value = ""; // Clear input
  });
};
