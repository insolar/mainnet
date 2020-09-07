export GO111MODULE ?= on
export GOSUMDB ?= sum.golang.org
export GOFLAGS ?= -mod=vendor
export GOBIN = ${PWD}/bin

BIN_DIR ?= bin
ARTIFACTS_DIR ?= .artifacts
INSOLAR = insolar
INSOLARD = insolard
BENCHMARK = benchmark
AIRDROP = airdrop
REQUESTER= requester

ALL_PACKAGES = ./...
GOBUILD ?= go build
GOTEST ?= go test

TEST_COUNT ?= 1
TEST_PARALLEL ?= 1
TEST_ARGS ?= -timeout 1200s
GOMAXPROCS ?= 0
TESTED_PACKAGES ?= $(shell go list ${ALL_PACKAGES})
COVERPROFILE ?= coverage.txt
BUILD_TAGS ?=

BUILD_NUMBER := $(TRAVIS_BUILD_NUMBER)
# skip git parsing commands if no git
ifneq ("$(wildcard ./.git)", "")
	BUILD_DATE ?= $(shell ./scripts/dev/git-date-time.sh -d)
	BUILD_TIME ?= $(shell ./scripts/dev/git-date-time.sh -t)
	BUILD_HASH ?= $(shell git rev-parse --short HEAD)
	BUILD_VERSION ?= $(shell git describe --tags)
endif
DOCKER_BASE_IMAGE_TAG ?= $(BUILD_VERSION)
DOCKER_IMAGE_TAG ?= $(DOCKER_BASE_IMAGE_TAG)

GOPATH ?= `go env GOPATH`
LDFLAGS += -X github.com/insolar/insolar/version.Version=${BUILD_VERSION}
LDFLAGS += -X github.com/insolar/insolar/version.BuildNumber=${BUILD_NUMBER}
LDFLAGS += -X github.com/insolar/insolar/version.BuildDate=${BUILD_DATE}
LDFLAGS += -X github.com/insolar/insolar/version.BuildTime=${BUILD_TIME}
LDFLAGS += -X github.com/insolar/insolar/version.GitHash=${BUILD_HASH}

INSGOCC=./bin/insgocc

.PHONY: all
all: submodule clean pre-build build ## cleanup, install deps, (re)generate all code and build all binaries

.PHONY: submodule
submodule: ## init git submodule
	git submodule init
	git submodule update

.PHONY: lint
lint: ## CI lint
	golangci-lint run

.PHONY: metalint
metalint: ## run gometalinter
	gometalinter --vendor $(ALL_PACKAGES)

.PHONY: clean
clean: ## run all cleanup tasks
	go clean $(ALL_PACKAGES)
	rm -f $(COVERPROFILE)
	rm -rf $(BIN_DIR)
	./insolar-scripts/insolard/launchnet.sh -l

.PHONY: install-build-tools
install-build-tools: ## install insolar tools for platform
	./scripts/build/ls-tools.go | xargs -tI % go install -v %

.PHONY: install-deps
install-deps: install-build-tools ## install dep and codegen tools

.PHONY: pre-build
pre-build: install-deps generate regen-builtin ## install dependencies, (re)generates all code

.PHONY: generate
generate: ## run go generate
	go generate -x $(ALL_PACKAGES)

.PHONY: ensure
ensure: ## does nothing (keep it until all direct calls of `make ensure` have will be removed)
	echo 'All dependencies are already in ./vendor! Run `make vendor` manually if needed'

.PHONY: vendor
vendor: ## update vendor dependencies
	rm -rf vendor
	go mod vendor

.PHONY: build
build: $(BIN_DIR) $(INSOLARD) $(INSOLAR) $(BENCHMARK)  $(AIRDROP) $(REQUESTER)## build all binaries
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

.PHONY: $(INSOLARD)
$(INSOLARD):
	$(GOBUILD) -o $(BIN_DIR)/$(INSOLARD) ${BUILD_TAGS} -ldflags "${LDFLAGS}" cmd/insolard/*.go

.PHONY: $(INSOLAR)
$(INSOLAR):
	$(GOBUILD) -o $(BIN_DIR)/$(INSOLAR) ${BUILD_TAGS} -ldflags "${LDFLAGS}" application/cmd/insolar/*.go

.PHONY: $(BENCHMARK)
$(BENCHMARK):
	$(GOBUILD) -o $(BIN_DIR)/$(BENCHMARK) -ldflags "${LDFLAGS}" application/cmd/benchmark/*.go

.PHONY: $(AIRDROP)
$(AIRDROP):
	$(GOBUILD) -o $(BIN_DIR)/$(AIRDROP) -ldflags "${LDFLAGS}" application/cmd/airdrop/*.go

.PHONY: $(REQUESTER)
$(REQUESTER):
	$(GOBUILD) -o $(BIN_DIR)/$(REQUESTER) application/cmd/requester/*.go

.PHONY: test_unit
test_unit: ## run all unit tests
	GOMAXPROCS=$(GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) -count=$(TEST_COUNT) -p=$(TEST_PARALLEL) $(TEST_ARGS) -json ./...| tee unit-test.log

.PHONY: functest
functest: ## run functest FUNCTEST_COUNT times
	GOMAXPROCS=$(GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) -test.v -p=$(TEST_PARALLEL) $(TEST_ARGS) -tags "functest" ./application/functest -count=$(TEST_COUNT)

.PNONY: functest_race
functest_race: ## run functest 10 times with -race flag
	make clean
	GOBUILD='go build -race' make build
	TEST_COUNT=10 make functest

.PHONY: test_func
test_func: functest ## alias for functest

.PHONY: test_slow
test_slow: ## run tests with slowtest tag
	CGO_ENABLED=1 GOMAXPROCS=$(GOMAXPROCS) \
		$(GOTEST) $(TEST_ARGS) -p=$(TEST_PARALLEL) -tags slowtest $(ALL_PACKAGES) -count=$(TEST_COUNT)

.PHONY: test
test: test_unit ## alias for test_unit

.PHONY: test_all
test_all: test_unit test_func test_slow ## run all tests (unit, func, slow)

.PHONY: test_with_coverage
test_with_coverage: $(ARTIFACTS_DIR) ## run unit tests with generation of coverage file
	CGO_ENABLED=1 $(GOTEST) $(TEST_ARGS) -tags coverage --coverprofile=$(ARTIFACTS_DIR)/cover.all --covermode=count $(TESTED_PACKAGES)
	@cat $(ARTIFACTS_DIR)/cover.all | ./scripts/dev/cover-filter.sh > $(COVERPROFILE)

.PHONY: test_with_coverage_fast
test_with_coverage_fast: ## ???
	CGO_ENABLED=1 $(GOTEST) $(TEST_ARGS) -tags coverage -count $(TEST_COUNT) --coverprofile=$(COVERPROFILE) --covermode=count $(ALL_PACKAGES)

$(ARTIFACTS_DIR):
	mkdir -p $(ARTIFACTS_DIR)

.PHONY: test-with-coverage
test-with-coverage: ## run unit tests with coverage, outputs json to stdout (CI)
	GOMAXPROCS=$(GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) $(TEST_ARGS) -json -v -count=$(TEST_COUNT) -p=$(TEST_PARALLEL) --coverprofile=$(COVERPROFILE) --covermode=count -tags 'coverage' $(ALL_PACKAGES)

.PHONY: regen-builtin
regen-builtin: ## regenerate builtin contracts code
	$(INSGOCC) regen-builtin -c application/builtin/contract -i github.com/insolar/mainnet/application/builtin/contract

.PHONY: build-track
build-track: ## get logs event tracker tool
	 GO111MODULE=off go get github.com/insolar/insolar/scripts/cmd/track

.PHONY: docker_base_build
docker_base_build: ## build base image with source dependencies and compiled binaries
	docker build -t insolar/mainnet-base:$(DOCKER_BASE_IMAGE_TAG) \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg BUILD_NUMBER="$(BUILD_NUMBER)" \
		--build-arg BUILD_HASH="$(BUILD_HASH)" \
		--build-arg BUILD_VERSION="$(BUILD_VERSION)" \
		-f ./scripts/kube/bootstrap/Dockerfile .
	docker tag insolar/mainnet-base:$(DOCKER_BASE_IMAGE_TAG) insolar/mainnet-base:latest
	docker images "insolar/mainnet-base"

.PHONY: docker_build
docker_build: ## build image with binaries and files required for kubernetes deployment.
	docker build -t insolar/mainnet:$(DOCKER_IMAGE_TAG) \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg BUILD_NUMBER="$(BUILD_NUMBER)" \
		--build-arg BUILD_HASH="$(BUILD_HASH)" \
		--build-arg BUILD_VERSION="$(BUILD_VERSION)" \
		-f Dockerfile .
	docker images "insolar/mainnet"

.PHONY: docker_clean
docker_clean: ## removes intermediate docker image layers w/o tags (beware: it clean up space, but resets caches)
	docker image prune -f

.PHONY: help
help: ## display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
