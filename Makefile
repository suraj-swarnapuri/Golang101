.PHONY: clean
	clean:
		go clean -cache -testcache -modcache

.PHONY: gen
gen:
	# Auto-generate code
	protoc --proto_path=. --twirp_out=. --go_out=. rpc/twirptest/twirptest.proto
