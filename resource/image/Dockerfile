FROM golang:alpine

WORKDIR /app
COPY main.go /app
RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
