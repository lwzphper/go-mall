PROJECT_NAME := "go-mall"
PKG := "github.com/lwzphper/go-mall"

BFF_OUTPUT_NAME := "bff"
BFF_MAIN_FILE_PATH := "bff/main.go"
MEMBER_OUTPUT_NAME := "member"
MEMBER_MAIN_FILE_PATH := "server/member/main.go"

BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"

all: build

dep: ## Get the dependencies
	@go mod tidy

build: dep ## Build the binary file
	@go build -a -o dist/${BFF_OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${BFF_MAIN_FILE_PATH}
	@go build -a -o dist/${MEMBER_OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MEMBER_MAIN_FILE_PATH}

linux: dep ## Build the binary file
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${BFF_OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${BFF_MAIN_FILE_PATH}
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${MEMBER_OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}'" ${MEMBER_MAIN_FILE_PATH}

init: dep ## Inital project
	@go run main.go init

run: dep ## Run Server
	@go run main.go start

clean: ## Remove previous build
	@go clean .
	@rm -f dist/${PROJECT_NAME}

install: ## Install depence go package
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen: ## Init Service
#	@protoc -I=. -I=common/pb --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} apps/*/pb/*.proto
#	@go fmt ./...

#	@protoc-go-inject-tag -input=apps/*/*.pb.go
	@./genProto.sh

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'