import {authService} from '../helpers/AuthServiceCheck.js';

let conn; // сохраняется между вызовами MessengerHandler

export default function MessengerHandler(element) {
    if (!element) { return }

    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    const pathParts = window.location.pathname.split('/');
    const user = pathParts[2]; // "kit2"
    const peer = pathParts[3]; // "jana"

    const timestamp = new Date().toLocaleString();


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
        sentMsg.innerText = `${msg.value} (${username}) ${timestamp}`;
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
                item.innerText = `${messages[i]} ( ${peer} ) ${timestamp}`;
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}