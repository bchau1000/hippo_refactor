FROM golang:1.18-alpine AS builder

ENV DB_HOST=host.docker.internal

# Replace with your MySQL login
ENV DB_USER=root
ENV DB_PASS=1234

# Create the working directory within the container
WORKDIR /tmp/build

# Copy all Go dependencies from our server files into the
# working directory
COPY go.mod .
COPY go.sum .

# Install Go dependencies from go.mod/go.sum
RUN go mod download

# Copy our server files into the container
COPY . .

# Build executable named 'deploy' from server files
RUN go build -o ./out/deploy .

# Create a smaller image to deploy onto
FROM alpine:3.9

ENV DB_HOST=host.docker.internal
ENV DB_USER=root
ENV DB_PASS=1234

WORKDIR /hippo-server
# Copy files built in builder image to our new smaller image
# into root directory /
COPY --from=builder /tmp/build/out/deploy /hippo-server

# Copy needed config.yaml files into /config
COPY --from=builder /tmp/build/config /hippo-server/config

# Run the server on 3000 within the container
EXPOSE 3000

# Run this command when the container starts
CMD ["/hippo-server/deploy"]