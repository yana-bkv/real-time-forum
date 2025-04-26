Right now I have one problem. After i submit login
my navigation stops working. But if you go back to 
website everything works. Everything else seems fine
with authorization.

Fow now I created authentication on the server with golang.
FRONTEND structure
Front using js. I put fetching functions in handlers and
html pages in views.

I will start to add posts table and adding joins to 
user table.

BACKEND structure
Controllers fetch json from page and create model
and the send it to database connect. 
In database connect we check if given info is 
valid and send it to sqlite to insert data to db.

1.03
I created basic crud for posts
create, find by id, find all, delete by id

i need to add joins to post and user
and represent data in one

tomorrow
now i need to add likes and comments and
need them to interact with each other
also maybe i need to make interaction
between user and comments differently


.........................
restructure front files/folders

add api fetch to handlers 
change /views folder

add tags 
show comments correctly

tests

docker 
documentation

мессенджер сделан на половину
надо добавить сохранение в бд
и готово!!! почти...
основная часть ))

* HTTP Request
* ↓
* Controller (UserController)
* ↓
* Repository (UserRepository)
* ↓
* Database (SQLite, etc.)


[Браузер клиента]
⬇ (сообщение через WebSocket)

[Hub (центральный менеджер всех подключений)]
⬇ (обрабатывает сообщение)

[Handler функции клиента]
⬇ (вызов сервиса для бизнес-логики)

[Service слой]
⬇ (сохраняет сообщение)

[Repository слой]
⬇ (вставка в базу данных)

[SQLite / Другая БД]

