#!/bin/bash

go build main.go
echo "\n main.go build "

pids=$(ps -eo comm,pid | awk '/^main/' | awk '{print $2}')

if [ -z "$pids" ]
then
    echo "empty"
else
    kill -9 $pids
    echo kill -9 $pids
fi

