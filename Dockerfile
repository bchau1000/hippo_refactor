FROM golang:1.18-alpine as build

# Create a working directory /temp
WORKDIR /temp

# Clone hippo-refactor into /temp
RUN git clone https://github.com/bchau1000/hippo_refactor.git /temp

# Since /temp/server already exists at this point this is equivalent
# to running 'cd /temp/server'
WORKDIR /temp/server

# Install Go dependencies defined by go.mod
RUN go mod download

# Build the Go server and create an executable: 'deploy' 
# we make a new directory solely for the executable, since
# we can only copy entire directories between stages
RUN go build -o ./out/deploy

# Create a new and smaller image w/ Alpine Linux
FROM alpine:3.15

# Create a working directory /server
WORKDIR /server

# Copy the executable we made in /temp/server into /server
# The path of this first copy is in reference to what I mentioned above
# 'deploy' is the only file in /temp/server/out and is therefore one of the few
# files that we will copy over
COPY --from=build /temp/server/out /server

# Run the executable we built in the first stage
CMD ["/server/deploy"]
