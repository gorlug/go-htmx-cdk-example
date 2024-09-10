#!/bin/sh

cd app
mkdir -p build/
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap
mv bootstrap build/
cp -r views build/
