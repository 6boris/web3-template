GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
COMPILE_TIME = $(shell date +"%Y%M%d")
CFLAGS += "\"$(COMPILE_TIME)\""
GIT_REVISION = $(shell git show -s --pretty=format:%h)
CFLAGS += "\"$(GIT_REVISION)\""

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find app -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: api
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
	       $(API_PROTO_FILES)
.PHONY: build

build:
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/interface/web3/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-interface-web3
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/service/data/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-service-data
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/service/demo/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-service-demo
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/job/datawatch/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-job-datawatch
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/job/databus/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-job-databus
	make build-app OUTPUT_FILE=bin/app INPUT_FILE=app/job/analysis/cmd/server/main.go APP_NAME=borisliu1/web3-template-app-job-analysis


build-app:
	docker build \
	  --build-arg VERSION=$(VERSION) \
	  --build-arg OUTPUT_FILE=$(OUTPUT_FILE) \
	  --build-arg INPUT_FILE=$(INPUT_FILE)  \
	  -t $(APP_NAME)  \
	  .
	docker tag $(APP_NAME):latest $(APP_NAME):$(COMPILE_TIME)-$(VERSION)
	docker push $(APP_NAME):latest
	docker push $(APP_NAME):$(COMPILE_TIME)-$(VERSION)

clean:
	docker compose stop
	docker compose rm
	rm -rf deploy/docker/mysql/data
	rm -rf deploy/docker/grafana/data
	rm -rf deploy/docker/prometheus/data
start:
	docker compose up -d
all:
	make api;
