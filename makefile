build:
	go build ./...

lint:
	golint -set_exit_status ./...
	staticcheck ./...

test:
	gotestsum --format testname -- -race ./...

install_tools:
	go install gotest.tools/gotestsum@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest

run:
	go run ./cmd/server
