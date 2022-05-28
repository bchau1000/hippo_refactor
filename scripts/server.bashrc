#!/bin/bash
DATE_NOW=$(date '+%Y/%m/%d %R:%S')
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
SERVER_DIR="${PROJECT_DIR}/server"
APP_DIR="${PROJECT_DIR}/app"

# Deploys the server locally and trails the server logs
function deploy() {
    kill $(ps aux | grep '/d/workspace/hippo_refactor/server/deploy' | awk '{print $2}') 2> /dev/null
    echo "${DATE_NOW} Starting server..." 
    cd ${SERVER_DIR} && rm deploy 2> /dev/null 2>&1 
    go build -o deploy . && nohup ./deploy > nohup.out 2> nohup.err < /dev/null &
    tail -n 0 -F "server.log"
}

function init() {
    python server.py -u $1 -p && exec bash
}

# Completely clears output in server.log
function clearLog() {
    echo "${DATE_NOW} Clearing output in server.log"
    truncate -s 0 ${SERVER_DIR}/server.log
}