# Use the official Golang image as the base image
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app
# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -mod=readonly -o swagger-binary ./swagger/main.go

# Use a lightweight image for the final image
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/swagger-binary /app/swagger-binary

# Set the working directory
WORKDIR /app

# Expose the port that the Swagger binary will run on
EXPOSE 8080

# Run the Swagger binary
CMD ["/app/swagger-binary"]
