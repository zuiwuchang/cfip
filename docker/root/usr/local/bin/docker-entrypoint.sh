#!/bin/bash
set -e
if [[ "$@" == "default-command" ]];then
    if [[ "$POST_URL" == "" ]];then
        exec cfip  -conf /data/v4.jsonnet
    else
        exec cfip  -conf /data/v4.jsonnet -url "$POST_URL"
    fi
else
    exec "$@"
fi