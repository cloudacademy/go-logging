all: logger customlogger filelogger logrus

logger:
	go run cmd/logger/main.go

filelogger:
	rm -f cloudacademy.log 2> /dev/null
	go run cmd/filelogger/main.go
	cat cloudacademy.log

customlogger:
	rm -f cloudacademy.log 2> /dev/null
	go run cmd/customlogger/main.go
	cat cloudacademy.log

logrus:
	rm -f cloudacademy.log 2> /dev/null
	go run cmd/logrus/main.go
	cat cloudacademy.log