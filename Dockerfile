FROM golang:latest

LABEL Author="Rakesh kumar Chanderki"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o hotspots

EXPOSE 8080

CMD ["./hotspots"]