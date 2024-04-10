# Support setting various labels on the final image
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.22-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git g++

# Get dependencies - will also be cached if we won't change go.mod/go.sum
COPY go.mod /prysm/
COPY go.sum /prysm/
RUN cd /prysm && go mod download

ADD . /prysm
RUN cd /prysm && go build -o=./build/bin/beacon-chain ./cmd/beacon-chain
# RUN cd /prysm && go build -o=./build/bin/validator ./cmd/validator

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache gcc musl-dev linux-headers git g++
COPY --from=builder /prysm/build/bin/beacon-chain /usr/local/bin/
COPY --from=builder /prysm/build/bin/validator /usr/local/bin/

EXPOSE 3500 4000 12000 13000
ENTRYPOINT ["beacon-chain"]
# ENTRYPOINT ["validator"]

# Add some metadata labels to help programatic image consumption
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
