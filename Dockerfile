FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-crud-api

EXPOSE 3000

CMD ["/go-crud-api"]
