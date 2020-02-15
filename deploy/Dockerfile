# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# enable go module
ENV GO111MODULE=on

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add ca-certificates git

# CompileDaemon for live reolad
RUN go get github.com/githubnemo/CompileDaemon

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
# * GOFLAGS=-mod=vendor to force `go build` to look into the `/vendor` folder, --
#  --> DO NOT use it if not using go mod vendor
# use go mod vendor command to copy your dependencies to /vendor in your app path.
ENV CGO_ENABLED=0
ENV GOFLAGS=-mod=vendor 

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/Mekawy5/chatserv

# # caching go dependencies, commented when using vendor
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Make wait-for.sh executable
RUN chmod +x wait-for.sh

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build \
    -installsuffix 'static' \
    -o /app ./cmd

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

