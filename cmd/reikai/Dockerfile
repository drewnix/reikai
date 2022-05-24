# Copyright 2022 the Reikai contributors.
# SPDX-License-Identifier: GPL-2.0

# syntax = docker/dockerfile:1

FROM bitnami/golang:1.18.2 as builder
WORKDIR /go/src/github.com/drewnix/reikai
COPY go.mod go.sum ./
COPY cmd cmd
ARG VERSION
# With the trick below, Go's build cache is kept between builds.
# https://github.com/golang/go/issues/27719#issuecomment-514747274
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -installsuffix cgo -ldflags \ 
       "-X github.com/drewnix/reikai/cmd/reikai/cmd.version=$VERSION" \
       ./cmd/reikai

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/drewnix/reikai/reikai /reikai
USER 1001
CMD ["/reikai"]
