#! /usr/bin/env bash
tag=$(git describe --tags|xargs)
moduleName=github.com/xieqiaoyu/xin-cli

BUILDARCH=amd64

if [[ ${OSTYPE} == darwin* ]]; then
    BUILDOS=darwin
elif [[ ${OSTYPE} == linux* ]]; then
    BUILDOS=linux
else
    BUILDOS=windows
    binName+='.exe'
fi

GO111MODULE=on packr2

CGO_ENABLED=0 GOOS=${BUILDOS} GOARCH=${BUILDARCH} go install -ldflags "-X '${moduleName}/metadata.Version=${tag}' -X '${moduleName}/metadata.Platform=${BUILDOS}_${BUILDARCH}' -s -w" .

packr2 clean
