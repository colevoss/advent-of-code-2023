.PHONY: run gen build-gen

run:
	go run $(CURDIR)/$(day)/*.go

gen:
	./generate $(day)

build-gen:
	go build -o generate ./gen/main.go
