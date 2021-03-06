include ./Makefile.common
# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
COMPILE_TIME = $(shell date +"%Y-%m-%d")
GIT_HASH = $(shell git log -n1 --format=format:"%h")
CFLAGS += $(COMPILE_TIME)-$(GIT_HASH)

binary = $(call get_binary_out,$(MOD),$(SURFIX))
TAG = cncamp-lilong-${MOD}-${SURFIX}-${CFLAGS}
TEST_LDFLAGS :=
EXTRA_TEST_ARGS :=
ifeq "$(DEBUG)" "1"
  EXTRA_TEST_ARGS= -v
endif
all:
	@mkdir -p ${BINARY_OUT_DIR}
	@make build MOD=mod2
release: build

build:
	GOOS=${PLATFORM} GOARCH=${ARCH} CGO_ENABLED=1 go build -mod=mod  ${GO_VERSION_FLAG} -v -o $(call get_binary_out,$(MOD),$(SURFIX)) $(call get_main,$(MOD))

push: build

test: build
	@echo "Running test"
	GOTEST='$(GOTEST)' TEST_LDFLAGS='$(TEST_LDFLAGS)' EXTRA_TEST_ARGS='$(EXTRA_TEST_ARGS)' TARGET_PATH='.' ./tools/test.sh
clean:
	go clean

run:
	./${binary} --config=file:///./config/mod8.ini --v=2 --logtostderr=true
docker-build: build
	cp ${binary} build/http-server
	cp -R config build/
	cd build && docker build . -t ${TAG}
docker-push: docker-build
	docker tag ${TAG} $(DOCKER_ACCOUNT)/${TAG}
	docker push $(DOCKER_ACCOUNT)/${TAG}
docker-run: docker-build
	docker run -d ${TAG}
version:
	@echo ${VERSION}

.PHONY: build test release clean version 
