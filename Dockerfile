# Use an official Go runtime as a parent image
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go API source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose a port that your Go application listens on
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
