test_static:
	golangci-lint run

test_unit: test_gen
	mkdir -p test
	go test -v ./...

AUTOGEN_FILES = \
	./internal/queryset/generator/test/autogenerated_models.go \
	./examples/comparison/gorm4/autogenerated_gorm4.go \
	./internal/queryset/generator/test/pkgimport/autogenerated_models.go

test_gen: gen
	@- $(foreach F,$(AUTOGEN_FILES), \
		go build $$(dirname $F)/*.go; \
	)

test: test_unit bench test_static

build:
	go build -o bin/goqueryset ./cmd/goqueryset/

bench:
	go test -bench=. -benchtime=1s -v -run=^$$ ./internal/queryset/generator/

gen: build
	@- $(foreach F,$(AUTOGEN_FILES), \
		go generate $$(dirname $F); \
	)