.PHONY: run gen

run:
	go run $(CURDIR)/$(day)/*.go

gen:
	./generate $(day)
