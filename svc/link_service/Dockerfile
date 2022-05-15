FROM golang:alpine AS builder

WORKDIR $GOROOT/src/github.com/Brijeshlakkad/delinkcious/svc/link_service

ADD ./main.go main.go
ADD ./service service

RUN go mod init link_service

# Fetch dependencies
RUN go get

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /link_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Brijesh Lakkad <brijeshlakkad22@gmail.com>
COPY --from=builder /link_service /app/link_service
EXPOSE 7070
ENTRYPOINT ["/app/link_service"]
