format:
	@gofumpt -l -w **/*.go

test-format:
	@./scripts/test-format.sh