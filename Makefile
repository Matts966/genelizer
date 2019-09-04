.PHONY: build install dep docker
YOUR_BINARY_NAME ?= generated
build:
	go test -count 1 -v ./...
	(cd generator && packr2 clean && packr2 build)
	go build -o $(YOUR_BINARY_NAME) --ldflags '-extldflags "-static"' ./cmd/genelize 
install: build
	mv $(YOUR_BINARY_NAME) $(GOPATH)/bin/
dep:
	go get -u github.com/gobuffalo/packr/v2/packr2@v2.6.0
docker:
	docker build -t genelizer . && docker run -it --rm genelizer
