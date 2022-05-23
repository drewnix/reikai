
IMPORT_PATH:= github.com/vmware-tanzu/kubeapps
GO = /usr/bin/env go
GOFMT = /usr/bin/env gofmt
IMAGE_TAG ?= dev-$(shell date +%FT%H-%M-%S-%Z)
VERSION ?= $$(git rev-parse HEAD)

IMG_MODIFIER ?=

default: all

reikai/%:
	DOCKER_BUILDKIT=1 docker build -t reikai/$*$(IMG_MODIFIER):$(IMAGE_TAG) --build-arg "VERSION=${VERSION}" -f cmd/$*/Dockerfile .

all: reikai/reikai
