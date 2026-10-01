FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /demo-app main.go

FROM alpine:3.19
WORKDIR /
COPY --from=builder /demo-app /demo-app
EXPOSE 8080

ENTRYPOINT ["/demo-app"]
