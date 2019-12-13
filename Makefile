help:	## display help
	# https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'¥

build:	## build notifier binary for Mac
	mkdir -p bin
	go build -o bin/notifier main.go

build-win:	## build notifier binary for Windows
	mkdir -p bin
	GOOS=windows GOARCH=amd64 go build -o bin/notifier.exe main.go

clean:	## cleanup binaries
	rm -rf ./bin/

.PHONY: help build build-win clean
