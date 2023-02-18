#!/bin/bash

FILES=$(gofumpt -l **/*.go)

if [ -z "${FILES}" ]; then
    echo "All good"
    exit 0
fi

echo "These files are not formatted, please run 'make format':"
echo "${FILES}"
exit 1