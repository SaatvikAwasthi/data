UNITTESTS=$(shell go list ./... |  grep -v /integrationTest | grep -v /server | grep -v /cmd | grep -v /mock)

compile:
	CGO_ENABLED=0 go build -o out/api ./cmd/api/*.go

start-server-locally:
	docker-compose -f ./docker-compose.yml up -d

end-server-locally:
	docker-compose -f ./docker-compose.yml down

fmt:
	go fmt ./...
	go vet ./...

dep:
	go mod tidy
	go mod download
	go mod vendor

test:
	go clean -testcache
	mkdir -p out/
	go test $(UNITTESTS) -coverprofile=out/coverage.out; \
    go tool cover -html=out/coverage.out -o out/coverage.html; \


generate:
	go generate ./...

 .PHONY: test fmt compile deps start-server-locally
