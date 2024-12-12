FROM golang:1.16
WORKDIR /app

COPY . .

RUN go mod init server || true
RUN go mod tidy 
RUN go build -o server .

CMD ["./server"]