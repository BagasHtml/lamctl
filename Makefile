.PHONY: build install uninstall clean

build:
	go build -o lamctl

install: build
	sudo cp lamctl /usr/local/bin/lamctl

uninstall:
	sudo rm -f /usr/local/bin/lamctl

clean:
	rm -f lamctl
