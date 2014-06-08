run:
	go run main.go
test:
	go test -cpu=1,2,4 ./...
cov:
	go test -coverprofile=coverage.out ./automata
	go test -cover ./regex
	go test -cover ./stack
	go test -cover ./lex
format:
	gofmt -s -l -w .
