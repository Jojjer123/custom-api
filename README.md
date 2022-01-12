# Custom-API

## Authors
Michael Mozooni & Gustav Rehnholm

## Introduction

custom-api is a Microservice for μONOS that receive RESTCONF messages from the user, convert it to gNMI and forward the message to the onos-config module.

A useful place to start with documentation is the ONOS Contributing guide or this projects prerequisite document.

- [Contributing to ONOS Project](https://docs.onosproject.org/developers/contributing/)
- [prerequisite](./docs/prerequisite.md)

## How to install
To get the google api files, run the following:
1. `mkdir -p google/api`
2. `curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto`
3. `curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto`

## How to run:
Run the following in the terminal:
1. `make clean`
2. `make generate`
3. `make run`
