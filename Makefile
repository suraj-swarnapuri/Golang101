.PHONY: test
test:
	@echo "==> running Go tests <=="
	CGO_ENABLED=1 go test -p 1 -v -race -short ./...

.PHONY: server
server:
	go build -gcflags='-N -l' -o bin/test-server ./cmd/8-Http/

.PHONY: run-server
run-server:
	./bin/test-server