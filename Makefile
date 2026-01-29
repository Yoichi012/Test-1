.PHONY: build run clean

build:
	go build -o bin/bot ./cmd/bot

run: build
	./bin/bot

clean:
	rm -rf bin/