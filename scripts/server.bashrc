#!/bin/bash
DATE_NOW=$(date '+%Y/%m/%d %R:%S')
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
SERVER_DIR="${PROJECT_DIR}/server"
APP_DIR="${PROJECT_DIR}/app"
IMAGE_NAME="hippo-server"

# Deploys the server locally and trails the server logs
function deploy() {
    kill $(ps aux | grep '/d/workspace/hippo_refactor/server/deploy' | awk '{print $2}') 2> /dev/null
    echo "${DATE_NOW} Starting server..." 
    cd ${SERVER_DIR} && rm deploy 2> /dev/null 2>&1 
    go build -o deploy . && nohup ./deploy > nohup.out 2> nohup.err < /dev/null &
    tail -n 0 -F "server.log"
}

# Deletes existing hippo-server images/containers and deploys a new one
function deployContainer() {
    CONTAINER_ID=$(docker ps | grep "${IMAGE_NAME}" | awk '{ print $1 }') &&

    echo "${DATE_NOW} Stopping existing ${IMAGE_NAME} container" &&
    docker stop ${CONTAINER_ID} 2> /dev/null;

    echo "${DATE_NOW} Removing existing ${IMAGE_NAME} container" &&
    docker rm ${CONTAINER_ID} 2> /dev/null;

    echo "${DATE_NOW} Removing existing ${IMAGE_NAME} image" &&
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep "${IMAGE_NAME}") 2> /dev/null;

    echo "${DATE_NOW} Building ${IMAGE_NAME} image" &&
    docker build --tag ${IMAGE_NAME} ${SERVER_DIR} &&

    echo "${DATE_NOW} Deploying ${IMAGE_NAME} to container" &&
    docker run -d -p 127.0.0.1:3000:3000/tcp ${IMAGE_NAME} &&
    docker ps;
}

# Completely clears output in server.log
function clearLog() {
    echo "${DATE_NOW} Clearing output in server.log"
    truncate -s 0 ${SERVER_DIR}/server.log
}