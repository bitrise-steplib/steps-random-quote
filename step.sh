#!/bin/bash

THIS_SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GO15VENDOREXPERIMENT="1"
go run "${THIS_SCRIPT_DIR}/step.go"
exit $?
