FROM golang:alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
COPY .env .



RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD [ "./app" ]