# simple-crud-streaming

Simple CRUDL (List as streaming HTTP/1.1 response) training example in hexagonal architecture with raw golang http server.
Made for [Geekbrains](https://gb.ru/) university course `Go Backend Level 1`.

Features:

- in-memory data storage
- three layers: adapters layer for api and database, core layyer, application layer
- basic auth middleware example

## Run

`go run ./cmd/reguser/main.go`

Insert some data:
`curl -u admin:admin -X POST -d '{"name":"user123"}' localhost:8000/create`
`curl -u admin:admin -X POST -d '{"name":"user5678"}' localhost:8000/create`

Search users:
`curl -u admin:admin 'localhost:8000/search?q=use'`

## App description for training
Create a simple CRUD user registration server (storage system - in memory).

The server must support operations:
- Create user profile as json
- Deleting a profile by a unique identifier
- Reading a profile by a unique identifier
- Getting all profiles by part of the name 

Additional requirements:
- Hexagonal architecture
- Graceful shutdown
- Basic auth middleware
- Streaming receiving a list of profiles, do not collect data in the handler 

## Live coding
Live video is [available here (Russian)](https://youtu.be/2D_2WWcpwEw)