#  Forum & Messenger App

A forum with private chats between users.
The backend is written in Go, and the frontend is embedded directly into the app (in the public/ folder).
Supports user registration, topic creation, replies, and private messaging via WebSocket.

---

## Tech Stack

	•	Go (Gorilla, WebSocket)
	•	SQLite
	•	Vanilla JavaScript
	•	Docker

---

## Project Structure
```
forum-messenger/
├── cmd/app/           # main.go — entry point
├── internal/
│   ├── auth/          # registration, login, tokens
│   ├── forum/         # topics and posts
│   ├── messenger/     # private chats via WebSocket
│   └── db/            # connection and migrations
├── public/            # frontend (HTML, CSS, JS)
├── go.mod             # dependencies
├── Dockerfile         # container build
└── README.md
```
##  How to Run

### Locally (without Docker)

```bash
go run main.go
```
### With Docker

```bash
docker build -t forum-app .
docker run -p 8080:8080 -p 3000:3000 forum-app
```

