# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
# For Go 1.3
FROM golang:1.3.3-onbuild

# For Go 1.6
FROM golang:1.6.2-onbuild

# Copy the package files to the container's workspace.
ADD . /go/src/github.com/fileparser

# Build the service command inside the container.
RUN go install github.com/fileparser

# Run service by default when the container starts.
ENTRYPOINT /go/bin/fileparser

# Document that the service listens on port 8080.
EXPOSE 8080