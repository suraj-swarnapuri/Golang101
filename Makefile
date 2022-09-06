.PHONY: test
test:
	@echo "==> running Go tests <=="
	CGO_ENABLED=1 go test -p 1 -v -race -short ./...