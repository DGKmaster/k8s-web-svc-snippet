# Build stage
FROM golang:1.18.2
WORKDIR /app
COPY svc/* ./
RUN go build -o app .

# Release stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 app ./
CMD ["./app"]  
