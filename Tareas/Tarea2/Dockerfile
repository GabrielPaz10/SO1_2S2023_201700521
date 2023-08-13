# Use the official Go image as a base
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

COPY *.go ./
# Build the application
RUN go build -o /Tarea2

# Port to listen on (on the container)
EXPOSE 3000

CMD [ "/Tarea2" ]