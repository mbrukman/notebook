#!/bin/bash -eu
#
# Copyright 2021 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

declare -r DOCKER_IMAGE="notebook-frontend"

build() {
  docker build \
    --build-arg UID="${UID}" \
    -t "${DOCKER_IMAGE}" \
    .
}

generic_run() {
  # The ${DOCKER_CMD} variable is intentionally not quoted so that it expands to
  # multiple distinct arguments, rather than a single string with spaces.
  docker run \
    -it \
    -p 8080:8080/tcp \
    -v "$(pwd)/../web/ui:/opt/notebook:rw" \
    "${DOCKER_IMAGE}" \
    ${DOCKER_CMD}
}

shell() {
  DOCKER_CMD="/bin/ash" generic_run
}

install() {
  DOCKER_CMD="yarn install" generic_run
}

run() {
  DOCKER_CMD="yarn run dev" generic_run
}

usage() {
  echo "Syntax: $(basename $0) [build | shell | install | run]" >&2
}

case "${1:-}" in
  build | shell | install | run)
    "$1"
    ;;

  "")
    echo "Command is required." >&2
    usage
    exit 1
    ;;

  *)
    echo "Unrecognized command: ${1:-}" >&2
    usage
    exit 2
esac
