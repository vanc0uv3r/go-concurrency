ARCH ?= arm64
OS ?= darwin
PROJECT_PATH=$(realpath ../)
BUILD_PATH=build

GO_ENV = GOARCH=${ARCH} GOOS=${OS}
BUILD_OUTPUT = ${PROJECT_PATH}/${BUILD_PATH}
BUILD_MSG = @echo "Built Done: ${BUILD_OUTPUT}/$(@F)"
START_MSG = @echo "Starting... ${BUILD_OUTPUT}/$(@F)"
RUN_MSG = @echo "Dev Running... ${BUILD_OUTPUT}/$(@F)"

# Targets

$(BUILD_PATH):
	mkdir -p ${PROJECT_PATH}/$(BUILD_PATH)


build_%: $(BUILD_PATH)
	$(GO_ENV) go build -o ${BUILD_OUTPUT}/$* ${PROJECT_PATH}/cmd/$*
	$(BUILD_MSG)

start_%:
	${BUILD_OUTPUT}/$*
	$(START_MSG)

run_%:
	go run ${PROJECT_PATH}/cmd/$*
	$(RUN_MSG)
