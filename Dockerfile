FROM golang:1.10-alpine as builder

# Install git so go get will work.
RUN apk add -U --no-cache git

# Pre-download listed dependencies to take
# advantage of Docker cache.
RUN mkdir -p /go/src/github.com/RiiConnect24/WUS-Go
WORKDIR /go/src/github.com/RiiConnect24/WUS-Go
COPY get.sh /go/src/github.com/RiiConnect24/WUS-Go
RUN sh get.sh

# Copy the entire WUS-Go source into builder's source.
COPY . .
RUN go get ./...

# Build to name "app".
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# Create a new image so we can have a smaller overall source iamge.
FROM disconnect24/docker-mail-runtime-base
WORKDIR /
COPY --from=builder /go/src/github.com/RiiConnect24/WUS-Go/app .

# Wait until there's an actual MySQL connection we can use to start.
CMD ["dockerize", "-wait", "tcp://database:3306", "-timeout", "60s", "/app"]