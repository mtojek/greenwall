build: go-get cc install test

go-get:
	go get golang.org/x/tools/cmd/goimports

cc:
	gofmt -s -w middleware
	goimports -w middleware

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	go vet -x ./...

dist: build
	zip -j dist.zip $(GOPATH)/bin/greenwall && zip -g -r dist.zip frontend

# TODO go lint

