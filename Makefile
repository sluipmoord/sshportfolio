

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./tests/...

## test/cover: run all tests and display coverage
.PHONY: test-cover
test-cover:
	go test -v -race -buildvcs -coverprofile=/tests/coverage.out ./tests/...
	go tool cover -html=/tests/coverage.out

## develop: run all tests and start the server
.PHONY: develop
develop:
	air -c .air.toml
	@echo 'Development server started'

## cli: run the CLI
.PHONY: cli
cli:
	go run ./cmd/cli/main.go
	@echo 'CLI started'

## build: build the server
.PHONY: build
build:
	go build -o ./bin/server ./cmd/ssh/main.go
	@echo 'Server built'

## build/cli: build the CLI
.PHONY: build/cli
build/cli:
	go build -o ./bin/cli ./cmd/cli/main.go
	@echo 'CLI built'

## run: run the server
.PHONY: run
run: build
	./bin/server
	@echo 'Server started'

## run/cli: run the CLI
.PHONY: run/cli
run/cli: build/cli
	./bin/cli
	@echo 'CLI started'

## clean: remove build artifacts
.PHONY: clean
clean:
	rm -rf ./bin
	rm -rf ./coverage.out
	rm -rf ./tests/coverage.out
	@echo 'Cleaned up build artifacts'

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code