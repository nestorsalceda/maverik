GODEPS = \
		github.com/gorilla/handlers \
		github.com/gorilla/mux \
		periph.io/x/periph \
		github.com/vektra/mockery/.../ \

TEST?=$(shell go list ./... | grep -v vendor)


all: test

test: mocks
	go test -i $(TEST)
	go test $(TEST)

mocks:
	mockery -all

build:
	go build

build-raspberry:
	GOOS=linux GOARCH=arm go build

clean:
	go clean
	rm -fr mocks

deps:
	go get -u -v $(GODEPS)
