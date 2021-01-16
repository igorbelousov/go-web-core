SHELL := /bin/bash

export PROJECT = Go-Web-Core


core:
	docker build \
	-f zarf/docker/dockerfile.go-web-core \
	-t go-web-core-amd64:1.0 \
	--build-arg VSF_REF=`git rev-parse HEAD` \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.

# ==============================================================================
# Running from within k8s/dev
kind-up:
	kind create cluster --image kindest/node:v1.19.4 --name starter-cluster --config zarf/k8s/dev/kind-config.yaml


kind-down:
	kind delete cluster --name starter-cluster

kind-load:
	kind load docker-image go-web-core-amd64:1.0 --name starter-cluster

kind-services:
	kustomize build zarf/k8s/dev | kubectl apply -f -


kind-status:
	kubectl get nodes
	kubectl get pods --watch


kind-status-full:
	kubectl describe pod -lapp=go-web-core

kind-update: core
	kind load docker-image go-web-core-amd64:1.0 --name starter-cluster
	kubectl delete pods -lapp=go-web-core


kind-logs:
	kubectl logs -lapp=go-web-core --all-containers=true -f
# ==============================================================================

run:
	go run ./cmd/app/main.go


runadmin:
	go run ./cmd/admin/main.go