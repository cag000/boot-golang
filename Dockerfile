FROM golang:1.13.12-alpine AS builder

# Copy project to working directory
RUN mkdir /app
COPY . /app/
WORKDIR /app/

RUN go build -o main *.go

# Run the binary.
ENTRYPOINT ["/app/main"]
