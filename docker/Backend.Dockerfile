FROM golang:1.13

LABEL Author="Rakesh kumar Chanderki"

WORKDIR /backend/main-framework

ADD ./backend/main-framework .

RUN go mod download

RUN go build -o main-framework

EXPOSE 8080

CMD ["./main-framework"]