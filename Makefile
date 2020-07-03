DATE := `date -u +%Y-%m-%d\T%H:%M:%SZ`
RELEASE_TAG := $$(git describe --tags --exact-match 2>/dev/null)
WORKING_TAG := $$(git describe --dirty --always --tags)
VERSION := `bash -c "if [ -z $(RELEASE_TAG) ]; then echo $(WORKING_TAG); else echo $(RELEASE_TAG); fi"`
TARGET := peud
MAIN_EXEC := ./cmd/peud/main.go
SHARED_BUILD_ENV := GOARCH=amd64 GO111MODULE=on
SHARED_BUILD_CMD := go build --ldflags "-X main.date=${DATE} -X main.version=${VERSION}" -tags netgo -installsuffix netgo
.PHONY: all build test clean docs

all: build

test_all: test vet fmt

test:
	go test ./... -coverprofile cover.out

build: build_linux build_win build_mac

build_linux:
	$(SHARED_BUILD_ENV) GOOS=linux $(SHARED_BUILD_CMD) -o bin/${TARGET}-linux ${MAIN_EXEC}

build_win:
	$(SHARED_BUILD_ENV) GOOS=windows $(SHARED_BUILD_CMD) -o bin/${TARGET}.exe ${MAIN_EXEC}

build_mac:
	$(SHARED_BUILD_ENV) GOOS=darwin $(SHARED_BUILD_CMD) -o bin/${TARGET}-darwin ${MAIN_EXEC}

docs:
	apidoc --config docs/apidoc.json -i internal/handlers/ -o docs/api/ --watch

run: fmt vet
	go run -race main.go ${ARGS}

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

clean:
	rm -rf ./bin docs/api
