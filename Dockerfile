# Multi stage building strategy for reducing image size.
FROM golang:1.20.11-alpine3.18 AS build-env
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app

# Install each dependencies
COPY go.mod /app
COPY go.sum /app
RUN go mod download
RUN apk add --no-cache --virtual git gcc make build-base alpine-sdk

# COPY main module
COPY . /app

# Check and Build
RUN make build-linux

### If use TLS connection in container, add ca-certificates following command.
### > RUN apk add --no-cache ca-certificates
FROM gcr.io/distroless/base-debian10
COPY --from=build-env /app/main /
EXPOSE 80
ENTRYPOINT ["/main"]