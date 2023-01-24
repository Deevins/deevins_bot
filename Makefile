.PHONY: run
run:
	go run cmd/bot/main.go


.PHONY: build
build:
	go build -o deevins_bot cmd/deevins_bot/main.go