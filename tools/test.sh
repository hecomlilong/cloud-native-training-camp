ARCH=$(go env GOARCH)
BINARY_DIR=./cmd
BINARY_OUT_LINUX=${BINARY_DIR}/linux/${ARCH}/mod2
BINARY_OUT_MACOS=${BINARY_DIR}/macos/${ARCH}/mod2
BINARY_OUT_WIN=${BINARY_DIR}/win/${ARCH}/mod2
set -e
trap 'set +e; PIDS=$(jobs -p); [ -n "$PIDS" ] && kill -9 $PIDS' EXIT
if [[ "${OS}" == "Windows_NT" ]]; then
 BINARY_OUT_DIR=${BINARY_OUT_WIN}
 PLATFORM="windows"
else
    if [[ "$(uname)" == "Darwin" ]]; then
  BINARY_OUT_DIR=${BINARY_OUT_MACOS}
  PLATFORM="macos"
    else
  BINARY_OUT_DIR=${BINARY_OUT_LINUX}
  PLATFORM="linux"
    fi
fi

echo $PLATFORM
web_server=${BINARY_OUT_DIR}

$web_server --v 2 --logtostderr true &
SERVER_PID=$!

echo "web_server PID:${SERVER_PID}"
TARGET_GO_FILES="./..."
cmd="cd ${TARGET_PATH} && $GOTEST ${EXTRA_TEST_ARGS} ${TARGET_GO_FILES} -cover true -check.p true -check.timeout 4s"
echo ${cmd} |awk '{run=$0;system(run)}'
