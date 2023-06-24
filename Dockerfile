FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o app .

RUN chmod +x app

EXPOSE 8002

CMD ["./app"]