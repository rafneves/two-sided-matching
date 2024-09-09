

.PHONY: \
	audit \
	setup \
	lint \
	imports \
	fmt \
	test \
	test/cover

setup:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html

test:
	go test -v -race -buildvcs ./...

lint:
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000,-ST1003 ./...

imports:
	goimports -l -w .

fmt:
	gofmt -l -w -s .


audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...

test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out