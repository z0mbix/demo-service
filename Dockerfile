# Step 1: Build Stage
FROM golang:1.21 as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build/server /app/server
CMD ["/app/server"]
