FROM --platform=linux/x86_64 golang:1.20-alpine

WORKDIR /usr/src/app

COPY . .

RUN go mod download

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
