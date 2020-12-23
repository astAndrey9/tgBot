FROM golang:1.15.6-buster

RUN mkdir /app


ADD . /app

WORKDIR /app

RUN apt-get update
RUN apt-get install -y git
RUN go get github.com/Syfaro/telegram-bot-api
RUN go build -o main .

CMD ["/app/main"]