.PHONY: build-app
build-app:
	$(GO_BUILD) -o ./bin/app ./cmd/app
