#!/bin/sh

rm -rf ./providers
git clone --depth 1 https://github.com/go-acme/lego.git /tmp/lego
cp -r /tmp/lego/providers/dns ./providers

cp /tmp/lego/challenge/dns01/dns_challenge_manual.go ./dns01
cp /tmp/lego/challenge/dns01/domain.go ./dns01
cp /tmp/lego/LICENSE ./providers
cp /tmp/lego/go.sum ./

find ./providers/ -name "*.go" -type f -exec sed -i 's+"github.com/go-acme/lego/v4/challenge/dns01"+"github.com/entrustcorporation/dv/dns01"+g' {} +
find ./providers/ -name "*.go" -type f -exec sed -i 's+"github.com/go-acme/lego/v4/providers/dns/+"github.com/entrustcorporation/dv/providers/+g' {} +

# Make sure we depend on the same versions
echo "module github.com/entrustcorporation/dv" > go.mod
echo "require (" >> go.mod
grep '^\s' /tmp/lego/go.mod >> go.mod
echo ")" >> go.mod

rm -rf /tmp/lego

go mod tidy

# Ensure we use the latest Entrust API client and Public Suffix List
go get -u github.com/entrustcorporation/entrust
go get -u github.com/weppos/publicsuffix-go