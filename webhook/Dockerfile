# start from official go image
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# copy source code from your host to your image filesystem
COPY . .

# build the go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# minimal image for final stage
FROM alpine:latest

WORKDIR /root/

# copy binary from builder stage
COPY --from=builder /app/main .

# run the binary
CMD ["./main"]
