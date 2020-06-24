FROM golang:1.14-alpine

ENV CGO_ENABLED 1
ENV GO11MODULE on

# hadolint ignore=DL3018
RUN apk add -U --no-cache bash=5.0.17-r0 make=4.3-r0 git=2.26.2-r0 \
    gcc=9.3.0-r2 g++=9.3.0-r2

RUN mkdir /app
COPY go.* /app/
WORKDIR /app

RUN go mod download

ENTRYPOINT [ "/bin/bash", "-c"]
CMD ["/bin/bash"]

# Docker image for running tests. This image is not strictly required,
# and that requires CGO to be enabled, which in turn requires GCC and G++ to be installed.
