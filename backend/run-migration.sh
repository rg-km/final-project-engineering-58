#!/bin/sh
export $(cat ./.env | xargs)
go run migrations/main.go