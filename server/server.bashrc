#!/bin/bash
DATE_NOW=$(date '+%Y/%m/%d %R:%S')
SERVER_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# Deploys the server locally and trails the server logs
function deploy() {
    echo "${DATE_NOW} Starting server..."
    cd ${SERVER_DIR} && rm deploy > /dev/null 2>&1
    go build -o deploy . && nohup ./deploy & > /dev/null 2>&1
    tail -n 0 -F "server.log"
}

# Completely clears output in server.log
function clearLog() {
    echo "${DATE_NOW} Clearing output in server.log"
    truncate -s 0 ${SERVER_DIR}/server.log
}