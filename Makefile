.PHONY: build install dep docker
YOUR_BINARY_NAME ?= generated
build:
	go test -count 1 ./...
	(cd generator && packr2 clean && packr2 build)
	go mod tidy
	go build -o $(YOUR_BINARY_NAME) --ldflags '-extldflags "-static"' ./cmd/genelize
	@echo "=== built binary named '${YOUR_BINARY_NAME}'!"
install: build
	mv $(YOUR_BINARY_NAME) $(GOPATH)/bin/
	@echo "=== install binary named '${YOUR_BINARY_NAME}'!"
dep:
	go mod download
	go get -u github.com/gobuffalo/packr/v2/packr2@v2.6.0
docker:
	docker build -t genelizer . && docker run -it --rm genelizer
