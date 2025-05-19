import {authService} from '../helpers/AuthServiceCheck.js';
import {api} from "../api/messages.js";

let conn; // сохраняется между вызовами MessengerHandler
let isRendered = false;

export default async function MessengerHandler(element) {
    if (!element) { return }

    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    const pathParts = window.location.pathname.split('/');
    const user = pathParts[2]; // "kit2"
    const peer = pathParts[3]; // "jana"
    console.log(user, peer);

    const timestamp = new Date().toLocaleString();

    renderMessages(user,peer);

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("messengerForm").onsubmit = async function (e) {
        e.preventDefault();
        const username = await authService.getUsername();
        // console.log(conn, msg.value);
        if (!conn) return false;
        if (!msg.value) return false;

        conn.send(msg.value);
        const sentMsg = document.createElement("div"); // message that user sents
        if (user != peer) {
            sentMsg.innerHTML = `
            <p class="msg-sender"><span class="msg-title">${username}:</span> ${msg.value} <span class="msg-timestamp">(${timestamp})</span></p>
        `;
        }
        appendLog(sentMsg);

        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:8080/ws/" + user + "/" + peer);

        if (conn && conn.readyState === WebSocket.OPEN) {
            return; // уже есть соединение
        }

        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };

        conn.onmessage = function (evt) {
            console.log("Received message: ", evt.data)

            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerHTML = `
                    <p class="msg-receiver"><span class="msg-title">${peer}: </span>${messages[i]} <span class="msg-timestamp">(${timestamp})</span></p>
                `;
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}

// Render messages from database
async function renderMessages(user, peer) {
    if (isRendered) return; // если уже отрисовано — выходим

    const messengerHistory = document.getElementById("messengerHistory");
    const result = await api.getMsg(user, peer);

    result.forEach(msg => {
        const msgElement = document.createElement("div");
        msgElement.classList.add("msg-item");
       if (user === peer) {
           msgElement.innerHTML = `
            <p class="msg-receiver"><span class="msg-title">${msg.sender}:</span> ${msg.content} <span class="msg-timestamp">(${msg.timestamp})</span></p>
        `;
       } else if (msg.sender === user) {
           msgElement.innerHTML = `
            <p class="msg-sender"><span class="msg-title">${msg.sender}:</span> ${msg.content} <span class="msg-timestamp">(${msg.timestamp})</span></p>
        `;
       } else {
           msgElement.innerHTML = `
            <p class="msg-receiver"><span class="msg-title">${msg.sender}:</span> ${msg.content} <span class="msg-timestamp">(${msg.timestamp})</span></p>
        `;
       }

        messengerHistory.appendChild(msgElement);
    });

    isRendered = true;
}
