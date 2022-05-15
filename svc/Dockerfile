# Build stage
FROM golang:1.18.2
WORKDIR /app
COPY ./* ./
RUN go build -o app .

# Release stage
FROM debian:stable-slim
WORKDIR /app
COPY --from=0 app ./
CMD ["./app"]  
