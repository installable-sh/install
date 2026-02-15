FROM golang:1.24-alpine AS builder
ARG VERSION=0.0.0-dev
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X github.com/installable-sh/lib/version.V=${VERSION}" -o INSTALL .

FROM scratch
COPY --from=builder /build/INSTALL /usr/local/bin/INSTALL
ENTRYPOINT ["/usr/local/bin/INSTALL"]
