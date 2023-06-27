# Build stage
FROM golang:1.19.0-alpine as build

WORKDIR /app
COPY . .
EXPOSE 8002

RUN CGO_ENABLED=0 GOOS=linux go build -o bank .

FROM alpine
COPY --from=build /app/bank /
COPY .env /



CMD ["/bank"]