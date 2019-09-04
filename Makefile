.PHONY: build install
build:
	go test -count 1 -v ./...
	(cd generator && packr clean && packr)
	go build -o $(YOUR_BINARY_NAME) ./cmd/genelize
install: build
	mv $(YOUR_BINARY_NAME) $(GOPATH)/bin/
dep:
	go get github.com/gobuffalo/packr/packr@v2.6.0
	go get github.com/gobuffalo/packr@v2.6.0
