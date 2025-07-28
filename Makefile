IMAGE_TAG = web

build:
	docker build --tag $(IMAGE_TAG) .

integration_test: build
	docker-compose up -d --wait
	go test -v `go list ./... | grep /integration`
	docker-compose down
