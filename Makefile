run:
	go run main.go
test:
	go test ./automata
	go test ./regex
	go test ./stack
format:
	gofmt -s -l -w .
