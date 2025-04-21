const socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = (event) => {
    console.log("New message:", event.data);
};

socket.onopen = () => {
    socket.send("Привет от клиента!");
};