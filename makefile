all: run

build:
	@go build -o ./bin/

run: build
	@./bin/github_activity