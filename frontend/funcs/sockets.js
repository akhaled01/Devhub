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
    console.log("Received message:", e.data);
    let data = JSON.parse(e.data);
    addMessage(data.req_Content.msg_content, false, data.req_Content.sender);
  };

  return ws;
};
