build:
	@go build -o ./build/grr

test:
	@go test -short ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ; do \
		golint $$file ; \
	done

run: build
	@./build/grr

.PHONY: build test lint run
