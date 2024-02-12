# Start from a Go image
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o /my-app .


# Expose the port that the Go application listens on
EXPOSE 8080

# Set the command to run the Go application
CMD ["/my-app"]