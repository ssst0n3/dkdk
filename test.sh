#!/bin/bash
#go test -failfast -count=1 -parallel 1 -v ./...
# https://github.com/golang/go/issues/33038
docker pull hello-world:latest
docker tag hello-world:latest 127.0.0.1:14005/dkdk/hello-world:v2
docker push 127.0.0.1:14005/dkdk/hello-world:v2
for s in $(go list ./...); do if ! go test -failfast -v -p 1 $s; then break; fi; done
