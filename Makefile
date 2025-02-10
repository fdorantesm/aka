.PHONY: build

build:
	mkdir -p dist
	# Read the AKA_VERSION value from .env and pass it to the compiledVersion variable via ldflags
	VERSION=$(shell grep '^AKA_VERSION=' .env | cut -d '=' -f 2) && \
	go build -ldflags "-X 'aka/src/cmd.compiledVersion=$$VERSION'" -o dist/aka ./src/main.go

install:
	cp dist/aka ${HOME}/bin/aka