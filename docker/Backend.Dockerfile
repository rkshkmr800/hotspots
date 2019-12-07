FROM golang:1.13

LABEL Author="Rakesh kumar Chanderki"

WORKDIR /app

COPY --from=builder /go/src/github.com/rkshkmr800/hotspots/backend/main-framework .

RUN go mod download

RUN go build -o hotspots

EXPOSE 8080

CMD ["./main-framework"]