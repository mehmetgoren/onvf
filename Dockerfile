FROM golang:1.18.4

RUN apt-get install -y tzdata

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]