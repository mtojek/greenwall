build: go-get cc install test

clean:
	rm -f $(GOPATH)/bin/greenwall
	rm -f dist.zip

go-get:
	go get golang.org/x/tools/cmd/goimports

cc:
	gofmt -s -w middleware *.go
	goimports -w middleware *.go

install:
	go get -t -v ./...

test:
	go test -v ./...
	go test -race  -i ./...
	go vet -x ./...

dist: build
	rm -f dist.zip && zip -j dist.zip $(GOPATH)/bin/greenwall && zip -g -r dist.zip frontend
	[ -s dist.zip ]

