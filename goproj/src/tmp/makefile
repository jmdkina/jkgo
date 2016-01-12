
SHELL=/bin/bash

FILE=main.go timeuse.go 
FILE += netuse.go

ECHO=$(@) echo

all: $(FILE)
	go build $^

install: $(FILE)
	go install $^
