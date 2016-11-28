export GOPATH = $(PWD)
export CONFPATH = $(GOPATH)/conf
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: clean ex1 ex2 ex3

all: client ex1 ex2 ex3

ex1:
	go install ex1

ex2:
	go install ex2/master1
	go install ex2/slave1

ex3:
	go install ex3/master2
	go install ex3/slave2

client:
	go install client

clean:
	rm -rf bin pkg

