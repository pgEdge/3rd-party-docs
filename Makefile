BINARY = bin/pgdoc-converter
SRC_DIR ?= ""
OUT_DIR ?= ./docs
MKDOCS ?= ./mkdocs.yml
VERSION ?= ""

.PHONY: build test lint clean convert validate

build:
	cd builder && go build -o ../$(BINARY) .

test:
	cd builder && go test ./... -v

lint:
	cd builder && test -z "$$(gofmt -l .)" || (gofmt -d . && exit 1) && go vet ./...

clean:
	rm -f $(BINARY)

convert: build
	./$(BINARY) -src $(SRC_DIR) -out $(OUT_DIR) \
		-mkdocs $(MKDOCS) -version $(VERSION) -verbose

validate: build
	./$(BINARY) -src $(SRC_DIR) -out $(OUT_DIR) \
		-mkdocs $(MKDOCS) -version $(VERSION) \
		-validate -verbose
