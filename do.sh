#!/usr/bin/env bash

if [[ $1 = 'validate' ]]; then
  swagger validate ./swagger.yml
elif [[ $1 = 'generate-server' ]]; then
  swagger generate server -A oauth-study -f ./swagger.yml
#elif [[ $1 = 'generate-spec' ]]; then
#    swagger generate spec ./swagger.yml
elif [[ $1 = 'resolve-dependencies' ]]; then
    go get -u -f ./...
elif [[ $1 = 'wire' ]]; then
    go get -u -f ./...
else
  echo 'unknown operation'
fi
