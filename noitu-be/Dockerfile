FROM golang:1.23 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/server /app/server
COPY --from=builder /app/config.yaml /config.yaml

EXPOSE 8081
CMD ["/app/server"]
