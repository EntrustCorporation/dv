#!/bin/sh

rm -rf ./providers
git clone --depth 1 git@github.com:go-acme/lego.git /tmp/lego
cp -r /tmp/lego/providers/dns ./providers
cp /tmp/lego/challenge/dns01/dns_challenge_manual.go ./dns01
cp /tmp/lego/LICENSE ./providers
rm -rf /tmp/lego

find ./providers/ -name "*.go" -type f -exec sed -i 's+"github.com/go-acme/lego/v4/challenge/dns01"+"github.com/digitorus/dv/dns01"+g' {} +
find ./providers/ -name "*.go" -type f -exec sed -i 's+"github.com/go-acme/lego/v4/providers/dns/+"github.com/digitorus/dv/providers/+g' {} +

go get -u
go mod tidy