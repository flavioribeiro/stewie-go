#!/bin/bash

function install_deps {
    	GOPATH=`pwd` go test -i stewie
    	GOPATH=`PWD` go install stewie/app
}

function run {
   	GOPATH=`pwd` go test stewie
}

function start_server {
    	GOPATH=`pwd` ./bin/app --GREP_PURPOSES &
}

function stop_server {
	ps aux | grep GREP_PURPOSES | awk '{print $2}' | xargs kill -9 &> /dev/null
}

function clean_bin {
	rm -rf bin/*
}

function main {
	clean_bin
    	install_deps
    	start_server
    	run
    	stop_server
}

if [ -z "$1" ]; then
    main
else
    $@
fi
