FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/*.go ./cmd/
COPY pkg/ ./pkg/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-crud-api ./cmd

EXPOSE 3000

CMD ["/go-crud-api"]
