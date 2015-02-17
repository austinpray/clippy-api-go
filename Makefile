all: clippy

clippy: clippy.go handlers.go middleware.go Capabilities.go SyncRequest.go
	go build .

clean:
	rm -f clippy-api-go

run: clippy
	PORT=9001 ./clippy-api-go

fmt:
	go fmt *.go

install: clippy
	cp -f clippy /usr/local/bin/clippy

test: clippy
	go test .

coverage: clippy
	go test . -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

install_deps:
	go get
