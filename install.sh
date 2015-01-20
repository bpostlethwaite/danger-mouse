#!/bin/bash

LOCAL_CONFIG=danger.json
CONFIG=/etc/danger.json
COMMAND=danger
BINARY=/usr/local/bin/$COMMAND

# Build binary
go build

# Make binary accessible on system
if [ ! -f $BINARY ]; then
    sudo ln -s $(pwd)/danger-mouse $BINARY
fi
# Copy config file if not copied
if [ ! -f $CONFIG ]; then
    sudo cp $LOCAL_CONFIG $CONFIG
fi
