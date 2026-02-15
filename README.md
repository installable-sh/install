# INSTALL

Installation tool for [installable.sh](https://github.com/installable-sh).

## Installation

### Docker (recommended)

```bash
docker pull ghcr.io/installable-sh/install:v1
```

### From source

```bash
go install github.com/installable-sh/install@latest
```

## Usage

```bash
INSTALL [options] <url> [args...]
```

> **Note:** INSTALL is under development and will be available in a future release.
> It is intended for installation and setup tasks during Docker image builds.

## Development

```bash
# Build
make build

# Run tests
make test

# Run all CI checks
make ci
```

## License

Apache 2.0 - see [LICENSE](LICENSE)
