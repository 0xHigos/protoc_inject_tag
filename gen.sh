#!/bin/bash
set -x
SCRIPT=$(readlink -f "$0")
CURT_DIR=$(dirname "$SCRIPT")
docker run \
  --rm \
  --net host \
  -u root:root \
  -v "$CURT_DIR":/llmops-common:rw \
  -v "${GOPATH}/include:/usr/include" \
  127.0.0.1:5000/tmp/protogen:with-inject-tag \
  bash \
  -c \
  "set -x &&
  apk add protobuf-dev &&
  ./proto-gen.sh"


