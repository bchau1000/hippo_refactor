#!/bin/bash

# TODO
function readConfig() {
    echo 'Reading config...'
}

# Must be in hippo directory to deploy the server locally
function deploy_local() {
    echo 'Starting server...'
    go run . && tail -F "server.log"
}

# TODO
function deploy() {
    echo  'Starting server...'
}

function clearLog() {
    echo 'Clearing server.log'
    rm server.log
}