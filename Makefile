BUILD_TIME:=$(shell date "+%Y%m%d-%H%M%S")

VERSION=$(shell git describe --tags --always --dirty --dirty="-dev")
ARCH=$(shell go env GOARCH)
MAIN_FILE=main.go
BINARY_DIR:=cmd
BINARY_OUT_LINUX:=$(BINARY_DIR)/linux/$(ARCH)/
BINARY_OUT_MACOS:=$(BINARY_DIR)/macos/$(ARCH)/
BINARY_OUT_WIN:=$(BINARY_DIR)/win/$(ARCH)/
MOD=mod2

ifeq ($(OS),Windows_NT)
 BINARY_OUT_DIR=$(BINARY_OUT_WIN)
 PLATFORM="windows"
else
 ifeq ($(shell uname),Darwin)
  BINARY_OUT_DIR=$(BINARY_OUT_MACOS)
  PLATFORM="macos"
 else
  BINARY_OUT_DIR=$(BINARY_OUT_LINUX)
  PLATFORM="linux"
 endif
endif
define get_binary_out
${BINARY_OUT_DIR}$(1)
endef
define get_main
$(1)/$(MAIN_FILE)
endef

all:
	@mkdir -p ${BINARY_OUT_DIR}
	@make 2
release: build

build:
	GOOS=${PLATFORM} GOARCH=${ARCH} CGO_ENABLED=1 go build -mod=mod  ${GO_VERSION_FLAG} -v -o $(call get_binary_out,$(MOD)) $(call get_main,$(MOD))
2:
	MOD=mod2
	@make build

push: build

test:
	go test -v ./...
clean:
	go clean

version:
	@echo ${VERSION}

.PHONY: build test release clean version 
