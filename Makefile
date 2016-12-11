
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

# TODO go lint

