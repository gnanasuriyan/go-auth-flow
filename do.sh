#!/usr/bin/env bash

if [[ $1 = 'validate' ]]; then
  swagger validate ./swagger.yml
elif [[ $1 = 'generate-server' ]]; then
  swagger generate server -A go-auth-flow -f ./swagger.yml
elif [[ $1 = 'resolve-dependencies' ]]; then
    go get -u -f ./...
elif [[ $1 = 'wire' ]]; then
    go run github.com/google/wire/cmd/wire ./app
elif [[ $1 = 'setup-local-db' ]]; then
    if [ ! "$(docker ps -a | grep 'dev_db')" ]; then
      cd docker/dev
      docker-compose up -d
      cd ../../
    fi
#    https://github.com/pressly/goose
    goose --dir ./db/migrations mysql "dev_user:password@tcp(127.0.0.1:3306)/acme-user?parseTime=true" up
else
  echo 'unknown operation'
fi
