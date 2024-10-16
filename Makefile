export PATH := $(PATH):`go env GOPATH`/bin
export GO111MODULE=on
LDFLAGS := -s -w

all: env fmt build

build: router-server router-client

env:
	@go version

# compile assets into binary file
file:
	rm -rf ./assets/router-server/static/*
	rm -rf ./assets/router-client/static/*
	cp -rf ./web/router-server/dist/* ./assets/router-server/static
	cp -rf ./web/router-client/dist/* ./assets/router-client/static

fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .

gci:
	gci write -s standard -s default -s "prefix(github.com/fatedier/frp/)" ./

vet:
	go vet ./...

router-server:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/router-server ./cmd/router-server

router-client:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/router-client ./cmd/router-client

test: gotest

gotest:
	go test -v --cover ./assets/...
	go test -v --cover ./cmd/...
	go test -v --cover ./client/...
	go test -v --cover ./server/...
	go test -v --cover ./pkg/...

e2e:
	./hack/run-e2e.sh

e2e-trace:
	DEBUG=true LOG_LEVEL=trace ./hack/run-e2e.sh

e2e-compatibility-last-router-client:
	if [ ! -d "./lastversion" ]; then \
		TARGET_DIRNAME=lastversion ./hack/download.sh; \
	fi
	FRPC_PATH="`pwd`/lastversion/router-client" ./hack/run-e2e.sh
	rm -r ./lastversion

e2e-compatibility-last-router-server:
	if [ ! -d "./lastversion" ]; then \
		TARGET_DIRNAME=lastversion ./hack/download.sh; \
	fi
	FRPS_PATH="`pwd`/lastversion/router-server" ./hack/run-e2e.sh
	rm -r ./lastversion

alltest: vet gotest e2e
	
clean:
	rm -f ./bin/router-client
	rm -f ./bin/router-server
	rm -rf ./lastversion

router-server-images:
	docker build . --file build/server/Dockerfile --tag docker.io/zxs943023403/router-server:v0.0.1

router-client-images:
	docker build . --file build/client/Dockerfile --tag docker.io/zxs943023403/router-client:v0.0.1

image-all: router-server-images router-client-images

docker:
	make build
	make image-all
