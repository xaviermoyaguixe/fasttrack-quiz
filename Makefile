build:
	go build -o quiz-cli . 

run-server: build
	./quiz-cli start-server

run-client: build
	./quiz-cli start-client

test:
	go test ./... -v -cover

clean:
	rm -rf bin/*
