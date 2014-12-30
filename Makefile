
deps:
	@echo "--> Installing build dependencies"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

test:
	go list ./... | xargs -n1 go test

.PHONY: all cov deps integ test web web-push