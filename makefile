SHELL := /bin/bash

export PROJECT = Go-Web-Core


core:
	docker build \
	-f zarf/dockerfile.go-web-core \
	-t go-web-core-amd64:1.0 \
	--build-arg VSF_REF=`git rev-parse HEAD` \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.



run:
	go run ./cmd/app/main.go


runadmin:
	go run ./cmd/admin/main.go