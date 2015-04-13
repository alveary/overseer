# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ENV PORT 8080
ENV app_dir /go/src/github.com/alveary/overseer

RUN go get github.com/codegangsta/gin
RUN go get github.com/kr/godep

# Copy the local package files to the container's workspace.
ADD . $app_dir
WORKDIR $app_dir

# CMD ls -la
CMD gin --port $PORT --godep run

# Document that the service listens on port 8080.
EXPOSE 8080
