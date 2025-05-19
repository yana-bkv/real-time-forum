#  Forum & Messenger App

Форум с приватными чатами между пользователями. Бэкенд на Go, фронтенд встроен прямо в приложение (папка `public/`).  
Поддерживает регистрацию, создание тем, ответы и личные сообщения через WebSocket.

---

##  Стек технологий

- Go (Gorilla, WebSocket)
- SQLite
- Чистый JavaScript
- Docker

---

## 📁 Структура проекта
```
forum-messenger/
├── cmd/app/           # main.go — точка входа
├── internal/
│   ├── auth/          # регистрация, логин, токены
│   ├── forum/         # темы и посты
│   ├── messenger/     # чаты через WebSocket
│   └── db/            # соединение и миграции
├── public/            # фронтенд (HTML, CSS, JS)
├── go.mod             # зависимости
├── Dockerfile         # сборка контейнера
└── README.md
```
##  Как запустить

###  Локально (без Docker)

```bash
go run main.go
```
### С докер

```bash
docker build -t forum-app .
docker run -p 8080:8080 -p 3000:3000 forum-app
```

