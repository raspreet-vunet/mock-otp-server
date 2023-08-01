# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Raspreet Singh <raspreet@vunetsystems.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Set environment variables
ARG HTTP_PORT=8085
ARG DATA_DIR=/opt/mock-otp-server/data/
ENV HTTP_PORT=$HTTP_PORT
ENV DATA_DIR=$DATA_DIR

# This container exposes port to the outside world
EXPOSE $HTTP_PORT

# Run the binary program produced by `go install`
CMD ["./main"]
