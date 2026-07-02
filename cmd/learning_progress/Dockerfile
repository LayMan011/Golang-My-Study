FROM golang:1.26.4-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/progress/exe /app/cmd/learning_progress

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/progress/exe /app
CMD ["/app/exe"]