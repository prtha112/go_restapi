test:
	echo "hello test"
	go test -v --run xxxxxx
	go test -v ./

document-swagger:
	swagger init spec \
		--title "A Todo list application" \
		--description "From the todo list tutorial on goswagger.io" \
		--version 1.0.0 \
		--scheme http
	touch Document/swagger.yml
	swagger generate spec -o ./Document/swagger.yml

build:
	go mod download
	GOARCH=amd64
	go build -o /go/bin/app

build-docker:
	docker build -t go_restapi -f Dockerfile .

run:
	rm -f ./main
	go run main.go