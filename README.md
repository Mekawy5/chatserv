# Golang Simple Chat Application



This is a dockerized golang simple chat system contains applications and application-chats.

  - you can create applications in the system.
  - each application contains multible chats.
  - messages should be saved in database.
  - user can search in messages.

--

> TODO.

  - messages should sent to queue instead of being saved in db directly.
  - create workers to consume the queue.
  - workers should save messages in database and elasticsearch.
  - search messages should be implemented.
  - use redis key-val im-memory db to store chat number
