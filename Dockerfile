# Use the official Go image from the Docker Hub
FROM golang:1.17

# Enable Go modules
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go script
RUN go build -o hdfs-downloader

# Set the command to run when the container starts
CMD ["/app/hdfs-downloader"]
