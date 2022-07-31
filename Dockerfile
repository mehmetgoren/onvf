FROM golang:1.18.4

WORKDIR /usr/src/app

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]