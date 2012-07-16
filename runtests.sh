#!/bin/bash

function install_deps {
    GOPATH=`pwd` go test -i stewie
}

function run {
    GOPATH=`pwd` go test stewie
}

function main {
    install_deps
    run
}

if [ -z "$1" ]; then
    main
else
    $@
fi
