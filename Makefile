.PHONY: all clean

all: main

main: main.go
	go get .
	go build -o main main.go

clean:
	rm main