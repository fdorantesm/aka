.PHONY: build

build:
	mkdir -p dist
	# Read the version value from version.txt and pass it to the compiledVersion variable via ldflags
	VERSION=$(shell cat version.txt) && \
	go build -ldflags "-X 'aka/src/cmd.compiledVersion=$$VERSION'" -o dist/aka ./src/main.go

install:
	cp dist/aka ${HOME}/bin/aka