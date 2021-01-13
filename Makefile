build:
	go build -o main *.go

run:
	go run *.go

docker:
	docker build . -t cag000/golang-test:latest
