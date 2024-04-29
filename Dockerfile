# Build the application from source
FROM golang:1.22-alpine AS builder

RUN apk add --update make

RUN mkdir /server
WORKDIR /server

COPY go.mod go.sum Makefile ./

RUN make setup

ADD . /server

RUN CGO_ENABLED=0 GOOS=linux make build

# Run the tests in the container
FROM builder AS tester

RUN make test

# Production image, copy all the files and run
FROM golang:1.21-alpine AS runner

WORKDIR /server

COPY --from=builder /server ./

EXPOSE 8080

CMD ["./bin/server"]
