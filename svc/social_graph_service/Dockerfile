FROM golang:alpine AS builder

WORKDIR $GOROOT/src/github.com/Brijeshlakkad/delinkcious/svc/social_graph_service

ADD ./main.go main.go
ADD ./service service

RUN go mod init social_graph_service

# Fetch dependencies
RUN go get -d -v

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /social_graph_service -a -tags netgo -ldflags '-s -w' .

FROM scratch
MAINTAINER Brijesh Lakkad <brijeshlakkad22@gmail.com>
COPY --from=builder /social_graph_service /app/social_graph_service
EXPOSE 9090
ENTRYPOINT ["/app/social_graph_service"]
