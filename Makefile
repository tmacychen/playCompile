SRC=$(wildcard *.go)
TEST=$(wildcard *_test.go)
GO_TEST=go test -v $^

all: ${SRC} ${TEST}
	go test -v .

simpleLexer: simpleLexer.go simpleLexer_test.go
#	go test -v $^
	${GO_TEST} 

simpleCalculator: simpleCalculator.go simpleCalculator_test.go
	${GO_TEST} $^

