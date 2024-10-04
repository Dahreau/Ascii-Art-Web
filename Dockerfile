# Use the official Golang image as the base image
FROM golang:1.22.4-alpine3.20

LABEL description="ascii-art-web Alex Dylan Daro"

# Set the working directory inside the container
WORKDIR /Ascii-Art-Web

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o ascii-art-webserver .

# Expose the port that the server will run on
EXPOSE 8080

# Run the Go web server
CMD ["./ascii-art-webserver"]
