CHAINNET?=betanet
BINARY?=sojanoded
GOPATH?=$(shell go env GOPATH)
GOBIN?=$(GOPATH)/bin
NOW=$(shell date +'%Y-%m-%d_%T')
COMMIT:=$(shell git log -1 --format='%H')
VERSION:=$(shell cat version)
IMAGE_TAG?=latest
HTTPS_GIT := https://github.com/sojahub/sojanode.git
DOCKER ?= docker
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf

GOFLAGS:=""
GOTAGS:=ledger

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=sojahub \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=sojanoded \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=sojanoded \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)' -tags '$(GOTAGS)'

BINARIES=./cmd/sojanoded ./cmd/sojagen ./cmd/ebrelayer ./cmd/sojatest

all: lint install

build-config:
	echo $(CHAINNET)
	echo $(BUILD_FLAGS)

init:
	./scripts/init.sh

start:
	sojanoded start

lint-pre:
	@test -z $(gofmt -l .)
	@GOFLAGS=$(GOFLAGS) go mod verify

lint: lint-pre
	@golangci-lint run

lint-verbose: lint-pre
	@golangci-lint run -v --timeout=5m

install: go.sum
	GOFLAGS=$(GOFLAGS) go install $(BUILD_FLAGS) $(BINARIES)

build-sojad: go.sum
	GOFLAGS=$(GOFLAGS) go build  $(BUILD_FLAGS) ./cmd/sojanoded

clean:
	@rm -rf $(GOBIN)/soja*

coverage:
	@GOFLAGS=$(GOFLAGS) go test -v ./... -coverprofile=coverage.txt -covermode=atomic

tests:
	@GOFLAGS=$(GOFLAGS) go test -v -coverprofile .testCoverage.txt ./...

feature-tests:
	@GOFLAGS=$(GOFLAGS) go test -v ./test/bdd --godog.format=pretty --godog.random -race -coverprofile=.coverage.txt

run:
	GOFLAGS=$(GOFLAGS) go run ./cmd/sojanoded start

build-image:
	docker build -t sojahub/$(BINARY):$(IMAGE_TAG) -f ./cmd/$(BINARY)/Dockerfile .

run-image: build-image
	docker run sojahub/$(BINARY):$(IMAGE_TAG)

sh-image: build-image
	docker run -it sojahub/$(BINARY):$(IMAGE_TAG) sh

init-run:
	./scripts/init.sh && ./scripts/run.sh

init-run-noInstall:
	./scripts/init-noInstall.sh && ./scripts/run.sh

rollback:
	./scripts/rollback.sh

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.3
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen.sh
.PHONY: proto-gen

proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace $(protoImageName) \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;
.PHONY: proto-format

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-any:
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) sh ./scripts/protocgen-any.sh
.PHONY: proto-gen-any

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh
.PHONY: proto-swagger-gen

proto-lint:
	$(DOCKER_BUF) lint --error-format=json
.PHONY: proto-lint

proto-check-breaking:
	# we should turn this back on after our first release
	# $(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master
.PHONY: proto-check-breaking
