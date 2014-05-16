run:
	go run main.go
test:
	go test ./automata
	go test ./regex
	go test ./stack
cov:
	go test -cover ./automata
	go test -cover ./regex
	go test -cover ./stack
format:
	gofmt -s -l -w .
