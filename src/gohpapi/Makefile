PKGS 		?= $(shell go list ./... | /usr/bin/grep -v /vendor/)
PKGS_DELIM 	?= $(shell echo $(PKGS) | sed -e 's/ /,/g')
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: clean deps run build docker docker-run test coverage

clean:
		rm -rf target; \
		rm -f cover.out cover.html

deps:
		go get -d -t -v

simplify:
		gofmt -s -l -w $(SRC)

run: deps
		go run *.go

build: deps
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o target/app .

docker: build
		docker build -t gohpapi .

docker-run: docker
		docker run -it -p8080:8080 gohpapi

test: deps
		go test -v ./...

coverage: test
	$(shell [ -e cover.out ] && rm cover.out)
	go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -I {} bash -c {}
	echo "mode: count" > cover.out
	grep -h -v "^mode:" *.coverprofile >> "cover.out"
	rm *.coverprofile
	go tool cover -html=cover.out -o=cover.html
