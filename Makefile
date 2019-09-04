.PHONY: build install dep
YOUR_BINARY_NAME ?= generated
build:
	go test -count 1 -v ./...
	(cd generator && packr clean && packr)
	go build -o $(YOUR_BINARY_NAME) ./cmd/genelize
install: build
	mv $(YOUR_BINARY_NAME) $(GOPATH)/bin/
dep:
	go get github.com/gobuffalo/packr/packr@v1.30.1
