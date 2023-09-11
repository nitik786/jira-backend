# Use the official Go image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the application code into the container (including app and conf folders)
COPY app/ app/
COPY conf/ conf/
COPY tests/ tests/

# Copy go.mod and go.sum (if using Go modules)
COPY go.mod .


# Download Go module dependencies
RUN go mod download

# Install Git and Mercurial (required for go get)
RUN apt-get update && apt-get install -y git mercurial

#install
RUN go install github.com/revel/cmd/revel@latest

# Install the Revel framework
RUN go get github.com/revel/cmd/revel

# Ensure the $GOPATH/bin directory is in your PATH
ENV PATH="$PATH:$GOPATH/bin"

CMD ["revel", "run", "-a", "app"]

