PROJECT_NAME := "upload-to-s3"
JOB_NAME := "terraform-job-extractor"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# OS / Arch we will build our binaries for
OSARCH := "linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386"

.SHELLFLAGS = -c # Run commands in a -c flag
.SILENT: ; # no need for @
.ONESHELL: ; # recipes execute in same shell

default: build

all: test build

init: gen-mock

clean: ## Clean dev files
	go clean -i ./...
	rm -vf \
		"./${PROJECT_NAME}" \
		./coverage.* \
		./cpd.*

dep: ## Get DEV dependencies
	go get -v -u github.com/mitchellh/gox && \
	go get -v -u github.com/mattn/goveralls && \
	go get -v -u github.com/mibk/dupl && \
	go get -v -u github.com/client9/misspell/cmd/misspell && \
	go get -v -u github.com/golang/mock/gomock && \
	go install github.com/golang/mock/mockgen && \
	go install github.com/tcnksm/ghr && \
	go get github.com/securego/gosec/cmd/gosec && \
	go mod tidy

dependencies: dep ## Get ALL dependencies
	go mod download

tidy:  ## Execute tidy comand
	go mod tidy

run:
	@go run main.go --bucket adl-dev-prometheus-bkt --key prometheus/tfstate

build: tidy ## Build the binary file
	go build -i -v $(PKG_LIST)

build-out: tidy ## Build the binary file
	go build -i -v $(PKG_LIST) -o ./dist/$(PROJECT_NAME)_$(VERSION)

cross-build: tidy ## Build the app for multiple os/arch
	gox -osarch=$(OSARCH) -output "dist/{{.OS}}_{{.Arch}}/${JOB_NAME}"

install: build ## Build the binary file
	go install

lint: ## Execute lint
	golangci-lint run ./...

fmt: ## Formmat src code files
	go fmt ${PKG_LIST}

cpd: ## CPD
	dupl -t 200 -html >cpd.html

test: ## Execute test
	echo "go test ${PKG_LIST}"
	go test -i ${PKG_LIST} || exit 1
	echo ${PKG_LIST} | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

race: ## Run data race detector
	go test -race -short ${PKG_LIST}

bench: ## Run benchmarks
	go test -bench ${PKG_LIST}

msan: ## Run memory sanitizer
	go test -msan -short ${PKG_LIST}

misspell: ## One way of improving the accuracy of your writing is to spell things right.
	misspell -locale US  .

coverage: ## Generate global code coverage report
	./scripts/coverage.sh;

gen-mock: ## Execute mockgen generatio go mocks
	mockgen -destination ./extractor/mock/roundtripper.go -package mhttp net/http RoundTripper

security: ## Execute go sec security step
	gosec -tests ./...

help: ## Display this help screen
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	all \
	init \
	clean \
	dep \
	dependencies \
	tidy \
	build \
	build-out \
	cross-build \
	install \
	package-release \
	release \
	lint \
	fmt \
	cpd \
	race \
	bench \
	vet \
	misspell \
	coverage \
	gen-mock \
	security
