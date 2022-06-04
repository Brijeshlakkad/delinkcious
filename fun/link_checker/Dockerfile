# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:alpine AS builder

# Copy local code to the container image.
WORKDIR $GOROOT/src/github.com/Brijeshlakkad/delinkcious/fun/link_checker
ADD ./main.go main.go

RUN go mod init link_checker

# Fetch dependencies
RUN go get

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /link_checker

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM registry.cn-hangzhou.aliyuncs.com/knative-sample/alpine-sh:3.9
# RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /link_checker /app/link_checker

# Run the web service on container startup.
CMD ["/app/link_checker"]
