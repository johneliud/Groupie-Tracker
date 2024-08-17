# Use golang-alpine as the base image
FROM golang:1.22.6:alpine

# Additional image metadata
LABEL maintainer="github.com/johneliud"
LABEL version="1.0"
LABEL description="Image file for a web server that consists of receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information."

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to the container
COPY go.mod ./

# Install the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Expose port to the outside world
EXPOSE 8080

# Set the command to run the application
CMD ["./main"]