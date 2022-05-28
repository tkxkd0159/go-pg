TARGET ?= hfgo
GCC_MINOR ?= 3

.PHONY: build clean run

run :
	@echo "###### Run $(TARGET) ######\n"
	@go run ./cmd/$(TARGET)

build :
	go build -o bin/$(TARGET) ./cmd/$(TARGET)

clean :
	rm -rf ./bin/
