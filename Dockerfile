# docker build .
# docker tag 0 hectorgool/godocker:version1.0

# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/hectorgool/gosanto

# Setting up working directory
WORKDIR /go/src/github.com/hectorgool/gosanto

# Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore 

# Build the taskmanager command inside the container.
RUN go install github.com/hectorgool/gosanto

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/gosanto

# Service listens on port 8080.
EXPOSE 8080