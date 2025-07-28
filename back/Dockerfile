# Step 1: Modules caching
FROM golang:1.24.5-alpine3.21 as modules

COPY go.mod go.sum /modules/

WORKDIR /modules

RUN go mod download

# Step 2: Builder
FROM golang:1.24.5-alpine3.21 as builder

COPY --from=modules /go/pkg /go/pkg
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./ && \
    go build -o /bin/migration ./migration && \
    chmod +x /bin/app /bin/migration

# Step 3: Final
FROM alpine:3.21

COPY --from=builder /app/config /config
COPY --from=builder /bin/migration /migration
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/app"]