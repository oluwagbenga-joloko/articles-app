BUILDPATH=$(CURDIR)
GO=$(shell which go)

EXENAME=articles_app

test:
	@$(GO) test  -v ./tests -coverpkg ./...

build: 
	@$(GO) build -o $(EXENAME)

run:
	./$(EXENAME)

all: test build run