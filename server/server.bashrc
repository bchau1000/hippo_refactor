#!/bin/bash

function deploy() {
    go run . && tail -F "server.log"
}