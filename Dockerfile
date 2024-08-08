FROM golang:1.22.5-alpine AS builder

WORKDIR /app

ENV ENVIRONMENT="production"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /bin/ama ./cmd/ama/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /bin/ama .

EXPOSE ${PORT}
EXPOSE ${AMA_DATABASE_PORT}

ENTRYPOINT [ "./ama" ]