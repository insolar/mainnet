export GO111MODULE ?= on
export GOSUMDB ?= sum.golang.org
export GOFLAGS ?= -mod=vendor
export GOBIN = ${PWD}/bin

BIN_DIR ?= bin
ARTIFACTS_DIR ?= .artifacts
INSOLAR = insolar
INSOLARD = insolard
BENCHMARK = benchmark
REQUESTER= requester

ALL_PACKAGES = ./...
GOBUILD ?= go build
GOTEST ?= go test

TEST_COUNT ?= 1
FUNCTEST_COUNT ?= 1
TESTED_PACKAGES ?= $(shell go list ${ALL_PACKAGES})
COVERPROFILE ?= coverage.txt
TEST_ARGS ?= -timeout 1200s
BUILD_TAGS ?=

CI_GOMAXPROCS ?= 8
CI_TEST_ARGS ?= -p 4

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
LDFLAGS += -X github.com/insolar/mainnet/version.Version=${BUILD_VERSION}
LDFLAGS += -X github.com/insolar/mainnet/version.BuildNumber=${BUILD_NUMBER}
LDFLAGS += -X github.com/insolar/mainnet/version.BuildDate=${BUILD_DATE}
LDFLAGS += -X github.com/insolar/mainnet/version.BuildTime=${BUILD_TIME}
LDFLAGS += -X github.com/insolar/mainnet/version.GitHash=${BUILD_HASH}

INSGOCC=./bin/insgocc

.PHONY: all
all: submodule clean pre-build build ## cleanup, install deps, (re)generate all code and build all binaries

.PHONY: submodule
submodule: ## init git submodule
	git submodule init
	git submodule update

.PHONY: lint
lint: ci-lint ## alias for ci-lint

.PHONY: ci-lint
ci-lint: ## CI lint
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
build: $(BIN_DIR) $(INSOLARD) $(INSOLAR) $(BENCHMARK) $(REQUESTER)## build all binaries
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

.PHONY: $(REQUESTER)
$(REQUESTER):
	$(GOBUILD) -o $(BIN_DIR)/$(REQUESTER) application/cmd/requester/*.go

.PHONY: test_unit
test_unit: ## run all unit tests
	CGO_ENABLED=1 $(GOTEST) -count=1 $(TEST_ARGS) $(ALL_PACKAGES)

.PHONY: functest
functest: ## run functest FUNCTEST_COUNT times
	CGO_ENABLED=1 $(GOTEST) -test.v $(TEST_ARGS) -tags "appfunctest" ./application/functest -count=$(FUNCTEST_COUNT)

.PNONY: functest_race
functest_race: ## run functest 10 times with -race flag
	make clean
	GOBUILD='go build -race' make build
	FUNCTEST_COUNT=10 make functest

.PHONY: test_func
test_func: functest ## alias for functest

.PHONY: test_slow
test_slow: ## run tests with slowtest tag
	CGO_ENABLED=1 $(GOTEST) -count=1 $(TEST_ARGS) -tags slowtest ./...

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
	CGO_ENABLED=1 $(GOTEST) $(TEST_ARGS) -tags coverage -count 1 --coverprofile=$(COVERPROFILE) --covermode=count $(ALL_PACKAGES)

$(ARTIFACTS_DIR):
	mkdir -p $(ARTIFACTS_DIR)

.PHONY: ci-test-with-coverage
ci-test-with-coverage: ## run unit tests with coverage, outputs json to stdout (CI)
	GOMAXPROCS=$(CI_GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) $(CI_TEST_ARGS) $(TEST_ARGS) -json -v -count 1 --coverprofile=$(COVERPROFILE) --covermode=count -tags 'coverage' $(ALL_PACKAGES)

.PHONY: ci-test-unit
ci-test-unit: ## run unit tests 10 times and -race flag, redirects json output to file (CI)
	GOMAXPROCS=$(CI_GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) $(CI_TEST_ARGS) $(TEST_ARGS) -json -v $(ALL_PACKAGES) -race -count 10

.PHONY: ci-test-slow
ci-test-slow: ## run slow tests just once, redirects json output to file (CI)
	GOMAXPROCS=$(CI_GOMAXPROCS) CGO_ENABLED=1 \
		$(GOTEST) $(CI_TEST_ARGS) $(TEST_ARGS) -json -v -failfast -tags slowtest ./... -count=$(TEST_COUNT)

.PHONY: ci-test-slow-long
ci-test-slow-long: ## run slow tests with race and count
	CI_TEST_ARGS=" -race " \
	TEST_ARGS=" -timeout 180m  " \
	TEST_COUNT=50 \
		$(MAKE) ci-test-slow

.PHONY: ci-test-slow-nightly
ci-test-slow-nightly: ## run slow tests with race and count (nightly run)
	CI_TEST_ARGS=" -race " \
	TEST_ARGS=" -timeout 480m  " \
	TEST_COUNT=80 \
		$(MAKE) ci-test-slow

.PHONY: ci-test-func-base
ci-test-func-base: ## run functest, redirects json output to file (CI)
	# GOMAXPROCS=2, because we launch at least 5 insolard nodes in functest + 1 pulsar,
	# so try to be more honest with processors allocation.
	GOMAXPROCS=$(CI_GOMAXPROCS) CGO_ENABLED=1  \
		$(GOTEST) $(CI_TEST_ARGS) $(TEST_ARGS) -json -tags "functest bloattest" -v ./application/functest -count=$(FUNCTEST_COUNT) -failfast

.PHONY: ci-test-func
ci-test-func:  ## run functest 3 times
	FUNCTEST_COUNT=3 \
		$(MAKE) ci-test-func-base

.PHONY: ci-test-func-long
ci-test-func-long: ## run functest with race and a little count
	CI_TEST_ARGS=" -p 10 -race " \
	TEST_ARGS=" -timeout 300m " \
	FUNCTEST_COUNT=10 \
		$(MAKE) ci-test-func-base

.PHONY: ci-test-func-nightly
ci-test-func-nightly: ## run functest with large count and race
	CI_TEST_ARGS=" -p 10 -race " \
	TEST_ARGS=" -timeout 1200m " \
	FUNCTEST_COUNT=200 \
		$(MAKE) ci-test-func-base

.PHONY: regen-builtin
regen-builtin: ## regenerate builtin contracts code
	$(INSGOCC) regen-builtin -c application/builtin/contract -i github.com/insolar/mainnet/application/builtin/contract

.PHONY: build-track
build-track: ## get logs event tracker tool
	 GO111MODULE=off go get github.com/insolar/insolar/scripts/cmd/track

.PHONY: docker_base_build
docker_base_build: ## build base image with source dependencies and compiled binaries
	docker build -t mainnet-base:$(DOCKER_BASE_IMAGE_TAG) \
		--build-arg BUILD_DATE="$(BUILD_DATE)" \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg BUILD_NUMBER="$(BUILD_NUMBER)" \
		--build-arg BUILD_HASH="$(BUILD_HASH)" \
		--build-arg BUILD_VERSION="$(BUILD_VERSION)" \
		-f docker/Dockerfile .
	docker tag insolar-base:$(DOCKER_BASE_IMAGE_TAG) mainnet-base:latest
	docker images "mainnet-base"

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
help: ## display help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
