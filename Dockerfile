# The base go-image
FROM golang:1.15.6-alpine3.12 AS builder
 
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

# Run command as described:
# go build will build an executable file named server in the current directory
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/web-server
RUN go build -o /app/web-server . 

FROM alpine:3.12

RUN mkdir /app
COPY --from=builder /app/web-server /app/web-server
COPY --from=builder /app/config /config

ENTRYPOINT ["/app/web-server"]