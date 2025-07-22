FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-app ./cmd/api/main.go

EXPOSE 13000

CMD ["./api-app", "--configFile", "env/docker.env"]
