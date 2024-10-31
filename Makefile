TARGET_OS ?= linux
TARGET_ARCH ?= amd64

IMAGE_NAME := $(shell basename $(CURDIR) | tr '[:upper:]' '[:lower:]')
CONTAINER_ID := $(shell docker ps --filter "ancestor=$(IMAGE_NAME)" --format "{{.ID}}")
DOCKER_EXEC := docker exec -it $(CONTAINER_ID)

.PHONY: build run ssh open clean

build:
	GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go build -o cmd/generate_pdf/generate_pdf cmd/generate_pdf/main.go
	docker build -t $(IMAGE_NAME) .
clean:
	rm -f cmd/generate_pdf/generate_pdf
	docker rmi -f bizcard-generator
run: build
	docker run --rm -P $(IMAGE_NAME)
ssh:
	$(DOCKER_EXEC) /bin/sh
open:
	@PORT=$$(docker port $(CONTAINER_ID) | grep -Eo '[0-9]+$$' | head -n 1) \
	&& [ -n "$$PORT" ] && open http://localhost:$$PORT || echo "Container is not running or has no exposed port"
