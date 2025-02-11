
FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o quiz-cli .

FROM scratch
COPY --from=builder /app/quiz-cli /quiz-cli

ENTRYPOINT ["/quiz-cli"]
