# The base go-image
FROM golang:1.15.6-alpine3.12 AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
 
# Create a directory for the app
RUN mkdir /app
 
# Set working directory
WORKDIR /app

# Get dependancies - will also be cached if we won't change mod/sum
COPY go.mod . 
COPY go.sum .
RUN go mod download

# Copy all files from the current directory to the app directory
COPY . /app


# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /app/web-server .

FROM alpine:3.12

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group


COPY --from=builder /app/web-server /go/bin/web-server
COPY --from=builder /app/config /config

USER appuser:appuser

ENTRYPOINT ["/go/bin/web-server"]