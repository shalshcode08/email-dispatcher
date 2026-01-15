# Simple Makefile for email-dispatcher

BINARY := email-dispatcher

.PHONY: build run clean

build:
	go build -o $(BINARY) .

run:
	go run .

clean:
	rm -f $(BINARY)
