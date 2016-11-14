export GOPATH = $(PWD)
export CONFPATH = $(GOPATH)/conf
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: clean ex1 ex2

all: client ex1 ex2

ex1:
	go install ex1

ex2:
	go install ex2/master1
	go install ex2/slave1

client:
	go install client

clean:
	rm -rf bin pkg

