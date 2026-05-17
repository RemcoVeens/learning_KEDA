#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o msg_consumer consumer/main.go
