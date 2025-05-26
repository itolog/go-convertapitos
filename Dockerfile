# Assembly stage
FROM golang:latest AS builder
WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Building the application
RUN CGO_ENABLED=0 go build -o /app/server ./backend/cmd/

# Final stage
FROM alpine:latest
WORKDIR /app

# Setting CA certificates and time zones
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

# Copy the binary file from the build step
COPY --from=builder /app/server /app/server
# А потом скопируйте бинарный файл в final stage:
COPY --from=builder /app/migrate /app/migrate

# Copy the necessary directories and files
COPY --from=builder /app/frontend/dist /app/frontend/dist
COPY --from=builder /app/public /app/public
COPY --from=builder /app/docs /app/docs
#COPY --from=builder /app/.env /app/.env

# Set the working directory
WORKDIR /app