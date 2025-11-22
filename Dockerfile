FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /bin/api ./cmd/api

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /bin/api /bin/api
COPY config.yaml config.yaml
COPY internal/docs/openapi.yaml internal/docs/openapi.yaml
EXPOSE 8080
CMD ["/bin/api"]
