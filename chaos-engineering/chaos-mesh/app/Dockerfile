FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /chaos-app main.go

FROM alpine:3.19
WORKDIR /
COPY --from=builder /chaos-app /chaos-app
EXPOSE 8080
ENTRYPOINT ["/chaos-app"]
