export GO111MODULE=on

default: test style get tidy build run

.PHONY: get
get:
	go get -v

.PHONY: build
build: get
	go build -v

.PHONY: test
test:
	go test

.PHONY: style
style:
	golangci-lint run -c .golangci.slow.yml *.go

.PHONY: tidy
tidy:
	go mod tidy -v

.PHONY: run
run:
	source .token.sh && github-searcher --query "sarama language:go org:Shopify"
