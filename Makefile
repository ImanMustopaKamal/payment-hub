.PHONY: build clean start zip format

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ping functions/ping/main.go

start: build
	@sls offline

zip: build
	@zip -j -9 ./bin/ping.zip ./bin/ping

format:
	gofmt -s -w .