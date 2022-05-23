
IMPORT_PATH:= github.com/vmware-tanzu/kubeapps
GO = /usr/bin/env go
GOFMT = /usr/bin/env gofmt
IMAGE_TAG ?= dev-$(shell date +%FT%H-%M-%S-%Z)
VERSION ?= $$(git rev-parse HEAD)

IMG_MODIFIER ?=

default: all

all: 
	$(GO) run main.go