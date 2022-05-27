ARG NUCLIO_LABEL=latest
ARG NUCLIO_ARCH=amd64
ARG NUCLIO_BASE_IMAGE=alpine:3.11
ARG NUCLIO_ONBUILD_IMAGE=nuclio/handler-builder-golang-onbuild:0.4.0-amd64

# Supplies processor uhttpc, used for healthcheck
FROM nuclio/uhttpc:0.0.1-amd64 as uhttpc

# Builds source, supplies processor binary and handler plugin
FROM ${NUCLIO_ONBUILD_IMAGE} as builder

# From the base image
FROM ${NUCLIO_BASE_IMAGE}

WORKDIR $GOROOT/src/github.com/Brijeshlakkad/delinkcious/fun/link_checker

ADD ./link_checker.go link_checker.go

RUN go mod init link_checker

# Fetch dependencies
RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o /link_checker -a -tags netgo -ldflags '-s -w' .

# Copy required objects from the suppliers
COPY --from=builder /home/nuclio/bin/processor /usr/local/bin/processor
COPY --from=builder /home/nuclio/bin/handler.so /opt/nuclio/handler.so
COPY --from=uhttpc /home/nuclio/bin/uhttpc /usr/local/bin/uhttpc

# Readiness probe
HEALTHCHECK --interval=1s --timeout=3s CMD /usr/local/bin/uhttpc --url http://127.0.0.1:8082/ready || exit 1

# Run processor with configuration and platform configuration
CMD [ "processor" ]