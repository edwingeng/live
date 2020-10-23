#!/usr/bin/env bash

[[ "$TRACE" ]] && set -x
pushd `dirname "$0"` > /dev/null
trap __EXIT EXIT

colorful=false
tput setaf 7 > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    colorful=true
fi

function __EXIT() {
    popd > /dev/null
}

function printError() {
    $colorful && tput setaf 1
    >&2 echo "Error: $@"
    $colorful && tput setaf 7
}

function printImportantMessage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

function printUsage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

function processProtoFile() {
    local OLD_DIR=`pwd`; cd "$1"
    echo "Processing $1/${@:2}..."
    protoc -I=. "${@:2}" --gogo_out=paths=source_relative:.
    [[ $? -ne 0 ]] && exit 1
    for arg in "${@:2}"; do
        go run github.com/mailru/easyjson/easyjson -all -omit_empty "`basename $arg .proto`.pb.go"
        [[ $? -ne 0 ]] && exit 1
    done
    cd "$OLD_DIR"
}

processProtoFile . pb.proto
