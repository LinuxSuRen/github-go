build:
	CGO_ENABLE=0 go build -ldflags "-w -s" -o bin/github-go
test:
	go test ./...
run:
	go run main.go

fmt:
	go mod tidy
	go fmt ./...

copy: build
	sudo cp bin/github-go /usr/local/bin/github-go

pre-commit: fmt test build
install-pre-commit:
	cp .github/pre-commit .git/hooks/pre-commit
