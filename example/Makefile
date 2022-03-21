
GOCMD=go
GOGEN=$(GOCMD) generate
GOTEST=$(GOCMD) test
GODEPS=$(GOCMD) mod tidy

imit-gen:
	$(GODEPS)
	$(GOGEN) ./gen

test:
	$(GOTEST)
