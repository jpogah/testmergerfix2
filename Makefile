.PHONY: run-hello-web build-hello-web docker-build-hello-web docker-run-hello-web

run-hello-web:
	go run ./cmd/hello-web

build-hello-web:
	mkdir -p bin
	go build -o bin/hello-web ./cmd/hello-web

docker-build-hello-web:
	docker build -f build/hello-web/Dockerfile -t hello-web:local .

docker-run-hello-web: docker-build-hello-web
	docker run --rm -p 8080:8080 hello-web:local
