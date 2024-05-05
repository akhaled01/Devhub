import { NewChatMessage } from "./utils";

/**
 * Creates a new ws connection to the backend
 * @returns ws - a websocket connection to DevHub backend
 */
export const NewChatWS = () => {
  let ws = new WebSocket("ws://localhost:8080/ws");

  ws.onopen = () => {
    console.log("websocket Opening Successful");
  };

  ws.onclose = () => {
    console.log("websocket closing Successful");
  };

  ws.onmessage = (e) => {
    console.log("RECIEVED MESSAGE:", JSON.parse(e.data));
    let data = JSON.parse(e.data);
    if (data.type === "message") {
      NewChatMessage(data.req_Content.msg_content, false, data.req_Content.sender);
    } else if (data.type === "online_user_list") {
      AssembleOnlineUsers(data);
    } else if (data.type === "Open_chat_response") {
      const message_space = document.getElementById("message-space");
      let data = JSON.parse(e.data);
      data.req_Content.forEach(message_obj, () => {
        message_space.innerHTML += `
        <div class="${message_obj.sender === sessionStorage.getItem("username") ? "mself" : "m"}">
          <div class="self message">
            <div class="sender-info">
                <div class="sname">You</div>
                <div class="date">Sat May 04 2024</div>
              </div>
            <p>wd</p>
          </div>
        </div>
        `;
      })
    }
  };

  return ws;
};

export const SaveCurrentChatUser = async (username) => {
  console.log(username);
  sessionStorage.setItem("chat_user", username);
  window.location.reload()
}

/**
 *  Takes JSON data and assembles online users
 * @param {*} data 
 */
export const AssembleOnlineUsers = (data) => {
  const contact_div = document.getElementById("c-contacts")
  contact_div.innerHTML = "";
  console.log(data);
  data.req_Content.forEach(user_obj => {
    let user = user_obj.user;
    console.log(user_obj.is_online);
    if (user.username !== sessionStorage.getItem("username")) {
      const contactDiv = document.createElement('div');
      contactDiv.classList.add('contact');
      contactDiv.id = user.username;
  
      const nameDiv = document.createElement('div');
      nameDiv.classList.add('name');
      nameDiv.textContent = user.username;
  
      contactDiv.appendChild(nameDiv);
      contact_div.appendChild(contactDiv);
  
      document.getElementById(user.username).addEventListener("click", () => {
        SaveCurrentChatUser(user.username);
      });
    }
  });
}
