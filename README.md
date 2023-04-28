# Stations

[![Build](https://github.com/hidromatologia-v2/stations/actions/workflows/build.yaml/badge.svg)](https://github.com/hidromatologia-v2/stations/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/hidromatologia-v2/stations/branch/main/graph/badge.svg?token=TKF3Y8FJJ3)](https://codecov.io/gh/hidromatologia-v2/stations)
[![Go Report Card](https://goreportcard.com/badge/github.com/hidromatologia-v2/stations)](https://goreportcard.com/report/github.com/hidromatologia-v2/stations)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/hidromatologia-v2/stations)

Stations microservice API

## Documentation

| File                                             | Description                                  |
| ------------------------------------------------ | -------------------------------------------- |
| [docs/spec.openapi.yaml](docs/spec.openapi.yaml) | OpenAPI specification for this microservice. |

## Installation

### Docker

```shell
docker pull ghcr.io/hidromatologia-v2/stations:0.0.2
```

### Docker compose

```shell
docker compose -f ./docker-compose.dev.yaml up -d
```

### Binary

You can use the binary present in [Releases](https://github.com/hidromatologia-v2/stations/releases/latest). Or compile your own with.

```shell
go install github.com/hidromatologia-v2/stations@latest
```

## Coverage

| [![coverage](https://codecov.io/gh/hidromatologia-v2/stations/branch/main/graphs/sunburst.svg?token=TKF3Y8FJJ3)](https://app.codecov.io/gh/hidromatologia-v2/stations) | [![coverage](https://codecov.io/gh/hidromatologia-v2/stations/branch/main/graphs/tree.svg?token=TKF3Y8FJJ3)](https://app.codecov.io/gh/hidromatologia-v2/stations) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |

