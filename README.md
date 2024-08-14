<div align="center">
  <img src="https://github.com/user-attachments/assets/f86cffbd-c1e8-4367-9f08-55378decbfdf">
  
  This **AMA (Ask me Anything)** repository holds the WebServer of the AMA project.
</div>

# ⁉ AMA | Ask me Anything
The project consists on creations of **ROOMS** where you can share with your audience in your next AMA Meeting.  
In the Room, anyone can make questions and upvote others people questions.

It uses `WebSocket` to subscribe to a room and get it updated realtime and handle some problems of race conditions using `mutual exclusion`. 

## 🖥 FrontEnd
The frontend of this application was developed using `React` and it's avaiable in its own repository.  

- **[AMA FRONTEND](https://github.com/rodrigomolter/ama)**

## 📋 Pre-requirements

- [Go version 1.22.5](https://go.dev/doc/install)

## 🌲 Environment Variables
Make a copy of the `env.template` file and rename it to `.env`  
Change the variables inside `.env` with the credentials of your database

## 🏗️ Usage

### 🔄 
Currently, the application does not perform `migrations` automatically, meaning you need to run the command to create the tables in the database.

> If you are using Docker Compose, this step can be initiated after the container starts.

To run the migrations, simply use `go run`.
```sh
go run cmd/tools/terndotenv/main.go
```

### 🐋 Docker Compose
You can start application using `docker-compose`. This will run the `backend`, `frontend` and the `postgres database`
```sh
docker-compose up
```
Make sure you have this [ama-backend](https://github.com/rodrigomolter/ama-backend) repository and the [ama-frontend](https://github.com/rodrigomolter/ama) repository in the same folder level.
```
/
└── ama/
    └── Dockerfile
└── ama-backend/
    └── Dockerfile
    └── docker-compose.yaml
```

### ⚙ Go
Or you can start only the API with the `go run` command
```sh
go run cmd/ama/main.go
```

## 🔍 Postman 
You can validate and manually test the API using Postman. Check here the Postman collection

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/29423847-510aa0f8-b3a8-4851-beb8-213f69201c2c?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D29423847-510aa0f8-b3a8-4851-beb8-213f69201c2c%26entityType%3Dcollection%26workspaceId%3D634d18e6-9a9f-45a7-a562-69e352023655)

## 🙌 Support this project

If you want to support this project, leave a ⭐.

Happy coding! 🚀

___

Made with ❤️ by [Rodrigo Molter](https://www.linkedin.com/in/rodrigo-molter/).
