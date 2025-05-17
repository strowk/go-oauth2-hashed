#!/bin/bash

docker run --rm -p 5432:5432 --name pg-example -e POSTGRES_PASSWORD=mysecretpassword postgres:latest &

sleep 5

go run main.go

docker stop pg-example