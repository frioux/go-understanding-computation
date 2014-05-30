run:
	go run main.go
test:
	go test ./automata
	go test ./regex
	go test ./stack
	go test ./lex
cov:
	go test -cover ./automata
	go test -cover ./regex
	go test -cover ./stack
	go test -cover ./lex
format:
	gofmt -s -l -w .
