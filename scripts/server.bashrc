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

# Deletes existing hippo-server images/containers and deploys a new one
# If on linux, run deploy.sh directly using: sudo bash deploy.sh
function deployContainer() {
    ./deploy.sh
}

# Completely clears output in server.log
function clearLog() {
    echo "${DATE_NOW} Clearing output in server.log"
    truncate -s 0 ${SERVER_DIR}/server.log
}