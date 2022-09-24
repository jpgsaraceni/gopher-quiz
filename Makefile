all:
	make compile
	make run

compile:
	go build -o ./bin/quiz main.go

run:
	./bin/quiz
