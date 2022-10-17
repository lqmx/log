mkPath := $(abspath $(lastword $(MAKEFILE_LIST)))
mkDir := $(dir $(mkPath))

.DEFAULT_GOAL := all

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: fpush
fpush:
	@go fmt ./...
	@git add .
	@git commit -v --no-edit --amend
	@git push --force

.PHONY: push
push:
	@go fmt ./...
	@git add .
	@git commit -m 'ok'
	@git push

.PHONY: github
github:
	@eval `ssh-agent -s` && ssh-add ~/.ssh/github_rsa

