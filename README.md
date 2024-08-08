<div align="center">
  <img src="https://github.com/user-attachments/assets/f86cffbd-c1e8-4367-9f08-55378decbfdf">
  
  This **AMA (Ask me Anything)** repository holds the WebServer of the AMA project.
</div>

# ⁉ Ask me Anything
The project consists on creations of **ROOMS** where you can share with your audience in your next AMA Meeting.  
In the Room, anyone can make questions and upvote others people questions.

It uses `WebSocket` to subscribe to a room and get it updated realtime and handle some problems of race conditions using `mutual exclusion`. 

## 🚀 Getting Started

### 📋 Pre-requirements

- [Go version 1.22.5](https://go.dev/doc/install)

### 🌲 Environment Variables
Make a copy of the `env.template` file and rename it to `.env`  
Change the variables inside `.env` with the credentials of your database

### 🏗️ Usage
You can setup a `Postgres` container using `docker compose` or use a local database.
```sh
docker-compose up
```

Start the API with the `go run`
```sh
go run cmd/ama/main.go
```

## 🙌 Support this project

If you want to support this project, leave a ⭐.

Happy coding! 🚀

___

Made with ❤️ by [Rodrigo Molter](https://www.linkedin.com/in/rodrigo-molter/).
