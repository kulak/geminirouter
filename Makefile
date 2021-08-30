.PHONY: default test tag

default:
	cat Makefile

test:
	go test -v ./...

tag:
	git tag $(name)
	git push origin $(name)
