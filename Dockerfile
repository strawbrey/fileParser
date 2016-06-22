# For Go 1.4
FROM golang:1.4.2-onbuild

# Copy the package files to the container's workspace.
ADD . /go/src/fileparser

# Build the service command inside the container.
RUN go install fileparser

# Run service by default when the container starts.
ENTRYPOINT /go/bin/fileparser

# Document that the service listens on port 8080.
EXPOSE 8080