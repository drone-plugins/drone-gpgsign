# drone-gpgsign

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-gpgsign/status.svg)](http://cloud.drone.io/drone-plugins/drone-gpgsign)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/gpgsign.svg)](https://microbadger.com/images/plugins/gpgsign "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-gpgsign?status.svg)](http://godoc.org/github.com/drone-plugins/drone-gpgsign)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-gpgsign)](https://goreportcard.com/report/github.com/drone-plugins/drone-gpgsign)

Drone plugin to sign artifacts with [GnuPG](https://gnupg.org/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-gpgsign/).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-gpgsign
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/gpgsign .
```

## Usage

```console
docker run --rm \
  -e PLUGIN_key=$(base64 -i path/to/private.key) \
  -e PLUGIN_PASSPHRASE=p455w0rd \
  -e PLUGIN_FILES=dist/* \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/gpgsign
```
