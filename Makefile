.PHONY: *
all: speakeasy docs

docs:
	go generate ./...

speakeasy: check-speakeasy
	bash ./fix_openapi.sh
	speakeasy generate sdk --lang terraform -o . -s ./openapi.yaml
	bash ./fix_provider.sh

check-speakeasy:
	@command -v speakeasy >/dev/null 2>&1 || { echo >&2 "speakeasy CLI is not installed. Please install before continuing."; exit 1; }
