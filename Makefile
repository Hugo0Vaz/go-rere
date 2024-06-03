build:
	@go build -o ./build/grr

run: build
	@./build/grr
