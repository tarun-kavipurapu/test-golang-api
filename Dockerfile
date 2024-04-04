# Start from the latest Alpine version with Go installed
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install necessary dependencies for building and running the application
RUN apk add --no-cache git
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
