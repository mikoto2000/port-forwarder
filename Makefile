APP_NAME := port-forwarder
GOARCH_DARWIN_AMD := amd64
GOARCH_DARWIN_ARM := arm64
GOARCH_LINUX := amd64
GOARCH_WINDOWS := amd64
WINDOWS_BINARY_NAME := ${APP_NAME}-windows-${GOARCH_WINDOWS}.exe
LINUX_BINARY_NAME := ${APP_NAME}-linux-${GOARCH_LINUX}
DARWIN_AMD_BINARY_NAME := ${APP_NAME}-darwin-${GOARCH_DARWIN_AMD}
DARWIN_ARM_BINARY_NAME := ${APP_NAME}-darwin-${GOARCH_DARWIN_ARM}

GO_BIN := ${GOPATH}/bin
VERSION := 1.1.0
LD_FLAGS := "-s -w -X main.version=${VERSION}"

DEST := ./build

WATCH_SRC := ./main.go

### 開発関連
# 開発環境の都合で、個別にビルドできるようにしている
# (Linux コンテナ上でコーディングを行い、 Windows 上で実行することがあるため)
all: build
build: build/${APP_NAME}
build/${APP_NAME}: ${WATCH_SRC}
	go build -ldflags=${LD_FLAGS} -trimpath -o ./build/${APP_NAME}

build-all: build-windows build-linux build-darwin-arm build-darwin-amd

build-windows: build/${WINDOWS_BINARY_NAME}
build/${WINDOWS_BINARY_NAME}: ${WATCH_SRC}
	GOOS=windows GOARCH=${GOARCH_WINDOWS} go build -ldflags=${LD_FLAGS} -trimpath -o build/${WINDOWS_BINARY_NAME}

build-linux: build/${LINUX_BINARY_NAME}
build/${LINUX_BINARY_NAME}: ${WATCH_SRC}
	GOOS=linux GOARCH=${GOARCH_LINUX} go build -ldflags=${LD_FLAGS} -trimpath -o build/${LINUX_BINARY_NAME}

build-darwin-arm: build/${DARWIN_ARM_BINARY_NAME}
build/${DARWIN_ARM_BINARY_NAME}: ${WATCH_SRC}
	GOOS=darwin GOARCH=${GOARCH_DARWIN_ARM} go build -ldflags=${LD_FLAGS} -trimpath -o build/${DARWIN_ARM_BINARY_NAME}

build-darwin-amd: build/${DARWIN_AMD_BINARY_NAME}
build/${DARWIN_AMD_BINARY_NAME}: ${WATCH_SRC}
	GOOS=darwin GOARCH=${GOARCH_DARWIN_AMD} go build -ldflags=${LD_FLAGS} -trimpath -o build/${DARWIN_AMD_BINARY_NAME}

.PHONY: lint
lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -checks inherit,ST1003,ST1016 ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf build
