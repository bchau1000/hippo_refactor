#!/bin/bash
DATE_NOW=$(date '+%Y/%m/%d %R:%S')
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROJECT_DIR="$(dirname "${SCRIPT_DIR}")"
SERVER_DIR="${PROJECT_DIR}/server"
IMAGE_NAME="hippo-server"

# Deletes existing hippo-server images/containers and deploys a new one
function deployContainer() {
    CONTAINER_ID=$(docker ps | grep "${IMAGE_NAME}" | awk '{ print $1 }') &&

    echo "${DATE_NOW} Stopping existing ${IMAGE_NAME} container" &&
    docker stop ${CONTAINER_ID};

    echo "${DATE_NOW} Removing existing ${IMAGE_NAME} container" &&
    docker rm ${CONTAINER_ID};

    echo "${DATE_NOW} Removing existing ${IMAGE_NAME} image" &&
    docker rmi $(docker images --format '{{.Repository}}:{{.Tag}}' | grep "${IMAGE_NAME}") 2> /dev/null;

    echo "${DATE_NOW} Building ${IMAGE_NAME} image" &&
    docker build --tag ${IMAGE_NAME} ${SERVER_DIR} &&

    echo "${DATE_NOW} Deploying ${IMAGE_NAME} to container" &&
    docker run -d -p 127.0.0.1:$1:3000/tcp ${IMAGE_NAME} &&
    docker ps;
}

if [ -z "$1" ]
then
    echo "You must provide a port to bind to (eg. ./deploy.sh 80)"
else
    deployContainer $1
fi