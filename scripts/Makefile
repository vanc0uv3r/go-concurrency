ARCH ?= arm64
OS ?= darwin
PROJECT_PATH=$(realpath ../)
STORAGE_PATH=cmd/storage
BUILD_PATH=build
STORAGE_BINARY=storage
BINARY_NAME=storage

$(BUILD_PATH):
	mkdir -p ${PROJECT_PATH}/$(BUILD_PATH)

build_storage: $(BUILD_PATH)
	GOARCH=${ARCH} GOOS=${OS} go build -o ${PROJECT_PATH}/${BUILD_PATH}/${STORAGE_BINARY} ${PROJECT_PATH}/${STORAGE_PATH}
	@echo "Built Done: ${PROJECT_PATH}/$(BUILD_DIR)/$(BINARY_NAME)"

start_storage: 
	${PROJECT_PATH}/${BUILD_PATH}/${STORAGE_BINARY}
	@echo "Starting... ${PROJECT_PATH}/$(BUILD_DIR)/$(BINARY_NAME)"

run_storage: 
	go run ${PROJECT_PATH}/${STORAGE_PATH}
	@echo "Dev Running... ${PROJECT_PATH}/$(BUILD_DIR)/$(BINARY_NAME)"