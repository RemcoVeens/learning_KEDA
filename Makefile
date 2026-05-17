__default__:
	install

format:
	go fmt ./...
	staticcheck ./...

test:
	go test --cover ./...
	gosec ./...

install:
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
