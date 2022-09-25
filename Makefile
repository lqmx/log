mkPath := $(abspath $(lastword $(MAKEFILE_LIST)))
mkDir := $(dir $(mkPath))

.DEFAULT_GOAL := all

.PHONY: env
env:
	@cd script && ./init.sh requirement

.PHONY: proto
proto:
	@cd script && ./proto.sh gencode

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: push
push:
	@go fmt ./...
	@git add .
	@git commit -v --no-edit --amend
	@git push --force

.PHONY: github
github:
	@eval `ssh-agent -s` && ssh-add ~/.ssh/github_rsa
