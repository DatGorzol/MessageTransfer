# MessageTransfer / Сервис Обмена Сообщениями

WebSocket Message Transfer Service built in Go.

Сервис обмена сообщениями через WebSocket, написанный на Go.

Supports real-time messaging, authentication, async dispatching, and offline message delivery.

Поддерживает обмен сообщениями в реальном времени, авторизацию, асинхронную очередь и offline-доставку.

---

## Features / Возможности

Real-time WebSocket communication  
JSON-based command protocol  
Authentication middleware  
Async dispatcher queue  
Offline message storage  
Thread-safe client manager  

Обмен сообщениями через WebSocket  
JSON-протокол команд  
Middleware авторизации  
Асинхронная очередь dispatcher  
Offline-хранилище сообщений  
Потокобезопасный менеджер клиентов  

---

## Protocol / Протокол

All communication uses JSON commands.

Все взаимодействие происходит через JSON-команды.

---

## Auth / Авторизация

```json
{
  "type": "auth",
  "data": {
    "token": "valid-token"
  }
}
```
## Message / Сообщение

```json
{
  "type": "message",
  "data": {
    "to": "user123",
    "text": "Hello / Привет"
  }
}
```

## Ping
```json
{
  "type": "ping"
}
```
| Method / Метод | Path / Путь | Description / Описание                      |
| -------------- | ----------- | ------------------------------------------- |
| GET            | /health     | Health check / Проверка сервиса             |
| WS             | /ws         | WebSocket connection / WebSocket соединение |


## Connection / Подключение
```bash
ws://localhost:8080/ws?client_id=YOUR_ID
```

## Example Usage / Пример использования

Example JS client:
```bash
const ws = new WebSocket("ws://localhost:8080/ws?client_id=user1");

ws.onmessage = (e) => console.log(e.data);

ws.onopen = () => {

  // auth
  ws.send(JSON.stringify({
    type: "auth",
    data: { token: "valid-token" }
  }));

  // send message
  ws.send(JSON.stringify({
    type: "message",
    data: {
      to: "user2",
      text: "Hello from browser / Привет из браузера"
    }
  }));
};
```
## Architecture / Архитектура

Client Manager
Dispatcher Queue
Offline Store
Command Handlers

Менеджер клиентов
Очередь dispatcher
Offline-хранилище
Обработчики команд

## Future Improvements / Возможные улучшения

JWT authentication
Redis offline storage
Worker pool
Heartbeat system
Persistent sessions

JWT авторизация
Redis для offline сообщений
Worker pool
Heartbeat / keepalive
Постоянные сессии
