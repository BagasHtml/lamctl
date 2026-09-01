.PHONY: build install uninstall clean lint test check

build:
	go build -o lamctl

install: build
	sudo cp lamctl /usr/local/bin/lamctl

uninstall:
	sudo rm -f /usr/local/bin/lamctl

clean:
	rm -f lamctl

lint:
	golangci-lint run ./...

test:
	go test ./...

check: lint test
