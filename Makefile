.PHONY: build

default: build

build:
	go build -ldflags "-s -w" -trimpath -o super-waffle
