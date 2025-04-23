import {authService} from '../helpers/AuthServiceCheck.js';

let conn; // сохраняется между вызовами MessengerHandler

export default function MessengerHandler(element) {
    if (!element) return;

    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("messengerForm").onsubmit = function () {
        if (!conn) return false;
        if (!msg.value) return false;

        conn.send(msg.value);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        if (conn && conn.readyState === WebSocket.OPEN) {
            return; // уже есть соединение
        }

        conn = new WebSocket("ws://" + document.location.host + "/ws");

        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };

        conn.onmessage = async function (evt) {
            const username = await authService.getUsername();
            const timestamp = new Date().toLocaleString();
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = `${messages[i]} ( ${username}) ${timestamp}`;
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
}