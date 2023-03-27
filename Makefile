

HASH=$(shell git rev-parse --short HEAD | tr -d '\n')
DATE=$(shell date +%FT%T%z)
VERSION=$(shell git describe --tags | tr -d '\n')

MOD_PATH=github.com/raylax/rayx

LDFLAGS=-ldflags "-w -s -X $(MOD_PATH)/version.BuildVersion=$(VERSION) -X $(MOD_PATH)/version.BuildDate=$(DATE) -X $(MOD_PATH)/version.BuildHash=$(HASH)"

GOBUILD:=go build -trimpath $(LDFLAGS) -a -o
BIN:=bin
NANE:=rayx

build:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(BIN)/$(NANE)_darwin_arm64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_darwin_amd64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_linux_amd64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(BIN)/$(NANE)_windows_amd64.exe