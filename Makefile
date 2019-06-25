all: check tests run

tests:
	@go test ./src/lexer
	@go test ./src/ast
	@go test ./src/parser

run:
	@go run src/main.go

docs:
	@pandoc ./REPORT.md -t html5 -o report.pdf

zip:
	@zip workclass.zip -r img src REPORT.pdf Makefile
