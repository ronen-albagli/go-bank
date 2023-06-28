# Build stage
FROM golang:latest AS build

WORKDIR /app
COPY . .
EXPOSE 8002

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o bank .

FROM alpine
COPY --from=build /app/bank /



CMD ["/bank"]