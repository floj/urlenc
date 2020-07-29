#!/usr/bin/env bash
set -eu
script_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export CGO_ENABLED=0
go build -v -o "${script_dir}/urlenc" "${script_dir}"

