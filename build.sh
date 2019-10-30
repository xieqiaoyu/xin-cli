#! /usr/bin/env bash
tag=$(git describe --tags|xargs)
moduleName=github.com/xieqiaoyu/xin-cli

BUILDARCH=amd64

binName="xin-cli"

if [[ ${OSTYPE} == darwin* ]]; then
    BUILDOS=darwin
    echo "mac"
elif [[ ${OSTYPE} == linux* ]]; then
    BUILDOS=linux
else
    BUILDOS=windows
    binName+='.exe'
fi

GO111MODULE=on packr2

CGO_ENABLED=0 GOOS=${BUILDOS} GOARCH=${BUILDARCH} go build -ldflags "-X '${moduleName}/metadata.Version=${tag}' -X '${moduleName}/metadata.Platform=${BUILDOS}/${BUILDARCH}' -s -w" -o artifact/${binName} .

packr2 clean
