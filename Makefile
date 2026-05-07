.PHONY: build run test docker-build deploy

build:
	go build -o bin/server ./src/main.go

run: build
	./bin/server

test:
	go test ./... -v

docker-build:
	docker build -t demo-app:latest .

deploy:
	kubectl apply -f k8s/
