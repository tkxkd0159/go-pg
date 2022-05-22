TARGET ?= hfgo
GCC_MINOR ?= 3

.PHONY: build clean run

run :
	go run ./cmd/$(TARGET)

build :
	go build -o bin/$(TARGET) ./cmd/$(TARGET)

clean :
	rm -rf ./bin/
