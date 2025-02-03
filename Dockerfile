# Use an official Golang image as a base image
FROM golang:1.22.5

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 3001

# Command to run the application
CMD ["./main"]
