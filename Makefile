.PHONY: build clean start zip format

clean:
	@go clean
	@rm -rf ./bin

build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/contacts/create functions/contacts/create/main.go

start: build
	@sls offline

zip: build
	@zip -j -9 ./bin/contacts/create.zip ./bin/contacts/create

format:
	gofmt -s -w .