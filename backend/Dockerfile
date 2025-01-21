FROM golang:1.21 AS builder

WORKDIR /backend-build

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o api-server ./cmd/main.go

# Create a minimal image
# Make workspace for the final image
FROM alpine:latest AS monolithic

WORKDIR /usr/local/rest-api

RUN apk add --no-cache tzdata

# EXPOSE 8080

COPY --from=builder /backend-build/api-server /usr/local/rest-api

# Copy example config
COPY --from=builder /backend-build/configs/config.yml.example /usr/local/rest-api/config.yml
# Copy db migration files
COPY --from=builder /backend-build/db /usr/local/rest-api/db
# Copy start script
COPY --from=builder /backend-build/scripts /usr/local/rest-api/scripts
RUN chmod +x /usr/local/rest-api/scripts/startup/*

ENTRYPOINT ["/usr/local/rest-api/api-server"]