test:
	echo "hello test"
	go test -v --run xxxxxx
	go test -v ./

swagger:
	swag init --md ./ -o Document

build:
	go mod download
	GOARCH=amd64
	go build -o /go/bin/app

build-docker:
	docker build -t go_restapi -f Dockerfile .

run:
	rm -f ./main
	go run main.go