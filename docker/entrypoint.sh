#!/bin/sh
set -x
go mod tidy && reflex -s -r .go go run "$1"