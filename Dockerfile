# Use an official Go runtime as a parent image
FROM golang:1.22.3-alpine3.18

# Set the working directory in the container
WORKDIR /app

# Copy the go module files and download dependencies
COPY go.mod ./
# COPY go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]
