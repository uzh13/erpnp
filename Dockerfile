# builder
FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY . .
RUN apk add --no-cache git
RUN go build -o /out/erpnp ./cmd/erpnp


# runtime
FROM alpine:latest
COPY --from=builder /out/erpnp /usr/local/bin/erpnp
ENTRYPOINT ["/usr/local/bin/erpnp"]