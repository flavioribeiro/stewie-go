#!/bin/bash

function install_deps {
    go test -i stewie
}

function run {
    go test stewie
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
