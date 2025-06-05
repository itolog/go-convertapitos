# Assembly stage
FROM golang:latest AS builder
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application and migrations
RUN CGO_ENABLED=0 go build -o /app/server ./backend/cmd/
RUN CGO_ENABLED=0 go build -o /app/migrate ./backend/migrations

# Final stage
FROM alpine:latest
WORKDIR /app

# Install ca-certificates and tzdata
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

# Copy the binaries from the build step
COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrate /app/migrate

# Copy the necessary directories and files
COPY --from=builder /app/frontend/dist /app/frontend/dist
COPY --from=builder /app/public /app/public
COPY --from=builder /app/docs /app/docs
# Just for testing
#COPY --from=builder /app/.env.test /app/.env

# Set the working directory
WORKDIR /app