.PHONY: build install dep docker
YOUR_BINARY_NAME ?= generated
build:
	(cd generator && packr2 clean && packr2 build && echo "package packrd_test" > packrd/packrd_test.go)
	go test -count 1 ./...
	go mod tidy
	go build -o $(YOUR_BINARY_NAME) --ldflags '-extldflags "-static"' ./cmd/genelize
	(cd generator && packr2 clean)
	@echo "=== built binary named '${YOUR_BINARY_NAME}'!"
install: build
	mv $(YOUR_BINARY_NAME) $(GOPATH)/bin/
	@echo "=== install binary named '${YOUR_BINARY_NAME}'!"
dep:
	go mod download
	go get -u github.com/gobuffalo/packr/v2/packr2@v2.6.0
docker:
	docker build -t genelizer . && docker run -it --rm genelizer
