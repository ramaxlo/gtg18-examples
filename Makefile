export GOPATH = $(PWD)
export CONFPATH = $(GOPATH)/conf
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: clean ex1

all: client ex1

ex1:
	go install ex1

client:
	go install client

clean:
	rm -rf bin pkg

