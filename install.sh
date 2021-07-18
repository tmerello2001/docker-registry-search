#!/bin/bash

if [ "$1" = "--short" ]; then
    mv docker-registry-search /usr/local/bin/drs
else
    if [ "$1" = "-short" ]; then
        mv docker-registry-search /usr/local/bin/drs
    else
        mv docker-registry-search /usr/local/bin
    fi
fi