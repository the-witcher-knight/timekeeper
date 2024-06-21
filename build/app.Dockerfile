FROM golang:1.22.4-alpine3.20

RUN apk update

WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Install sqlboiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest && \
    go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# Install mockery
RUN go install github.com/vektra/mockery/v2@v2.43.2

# Install abigen
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest
