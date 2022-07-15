#!/bin/bash

if [ $# != 2 ]; then
    echo "not enough arguments"
    return
fi

function block() {
    while IFS= read -r line; do
        iptables -A INPUT -s "$line" -j DROP
    done < $1
}

function unblock() {
    while IFS= read -r line; do
        iptables -A INPUT -s "$line" -j ACCEPT
    done < $1
}

if [ -f "$2" ]; then
    if [ "$1" = "block" ]; then
        block "$2"
    elif [ "$1" = "unblock" ]; then
        unblock "$2"
    else
        echo "invalid command"
    fi
else
     echo "file not found"
fi
