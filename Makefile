GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
MAIN_NAME=main/main.go
BINARY_NAME=./chord.exe


build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_NAME)

test: clean build
	$(BINARY_NAME) -a localhost -p 8000 --ja 128.8.126.63 --jp 4170 --ts 3000 --tff 1000 --tcp 3000 -r 4

clean:
	rm -f $(BINARY_NAME)

default: test