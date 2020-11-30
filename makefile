.PHONY: build test release

build:
	go build
	rm -rf web/build
	cd web && npm run build

test:
	go test ./...

release:
	make build
	tar -zcvf console.tar.gz please copy.bat web/build