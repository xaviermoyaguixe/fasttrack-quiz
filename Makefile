build:
	go build -o quiz-cli . 

run-server: build
	./quiz-cli start-server

run-client: build
	./quiz-cli start-client

clean:
	rm -rf bin/*
