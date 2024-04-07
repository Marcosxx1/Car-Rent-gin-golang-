FROM golang:1.22.1

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire directory into the container
COPY . .

# Change directory to where the main.go is located
WORKDIR /usr/src/app/api/infra/http

# Build the Go application
RUN go build -v -o /usr/local/bin/app .

# Reset working directory back to the root of the project
WORKDIR /usr/src/app

# Command to run the application
CMD ["app"]
