FROM golang:1.18-alpine

WORKDIR "/app"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o pizzaria-backend cmd/main.go

CMD [ "./pizzaria-backend"]