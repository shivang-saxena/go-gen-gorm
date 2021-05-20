test_static:
	golangci-lint run

test_unit: test_gen
	mkdir -p test
	go test -v ./...

AUTOGEN_FILES = \
	./internal/queryset/generator/test/generated_models.go \
	./examples/comparison/gorm4/generated_gorm4.go \
	./internal/queryset/generator/test/pkgimport/generated_models.go

test_gen: gen
	@- $(foreach F,$(AUTOGEN_FILES), \
		go build $$(dirname $F)/*.go; \
	)

test: test_unit bench test_static

build:
	go build -o bin/go-gen-gorm ./cmd/go-gen-gorm/

bench:
	go test -bench=. -benchtime=1s -v -run=^$$ ./internal/queryset/generator/

gen: build
	@- $(foreach F,$(AUTOGEN_FILES), \
		go generate $$(dirname $F); \
	)
