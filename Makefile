
all: build

clean:
	go clean -i ./...

test:
	go test -tags ci -cover ./...

build: test
	go build ./...

update:
	go get -u ./...
