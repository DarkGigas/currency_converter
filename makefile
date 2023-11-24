.PHONY: all build run test clean

APP_NAME := currency_converter
GO := go
FIBER := $(GO) run main.go

all: build

build:
	@$(GO) build -o $(APP_NAME) .

run:
	@$(FIBER)

test:
	@$(GO) test ./...

clean:
	@rm -f $(APP_NAME)
