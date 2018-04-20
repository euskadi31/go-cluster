.PHONY: all clean fmt vet test

PACKAGES = $(shell go list ./...)

release:
	@echo "Release v$(version)"
	@git pull
	@git checkout master
	@git pull
	@git checkout develop
	@git flow release start $(version)
	@git flow release finish $(version) -p -m "Release v$(version)"
	@git checkout develop
	@echo "Release v$(version) finished."

all: test

clean:
	@go clean -i ./...

fmt:
	@go fmt $(PACKAGES)

vet:
	@go vet $(PACKAGES)

test:
	@go test -cover -covermode=count -coverprofile coverage.out ./...

cover: test
	@echo ""
	@go tool cover -func coverage.out

