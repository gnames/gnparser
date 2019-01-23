#!/bin/bash
# ensure that the specified version of protoc is installed in
# /tmp/proto$PROTO_VERSION/bin/protoc, which maybe cached

# bash scripts/protoc-install.sh 3.6.1


# make bash more robust.
set -beEux -o pipefail

if [ $# != 1 ] ; then
    echo "wrong # of args: $0 protoversion" >&2
    exit 1
fi

PROTO_VERSION="$1"
PROTO_FILE="protoc-${PROTO_VERSION}-linux-x86_64.zip"
PROTO_PATH="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/${PROTO_FILE}"
PROTO_DIR="/tmp/proto-${PROTO_VERSION}"

# Can't check for presence of directory as cache auto-creates it.
if [ ! -f "${PROTO_DIR}/bin/protoc" ]; then
  wget -O "${PROTO_FILE}" "${PROTO_PATH}" 
  unzip -o "${PROTO_FILE}" -d "${PROTO_DIR}"
  rm "${PROTO_FILE}"
  cp "${PROTO_DIR}/bin/protoc" /bin
fi
