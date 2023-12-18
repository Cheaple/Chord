GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
MAIN_NAME=main/main
BINARY_NAME=chord


build:
	$(GOBUILD) $(MAIN_NAME) -o $(BINARY_NAME) -v

test: build
	./chord -a localhost -p 8000 --ja 128.8.126.63 --jp 4170 --ts 3000 --tff 1000 --tcp 3000 -r 4

clean:
	rm -f $(BINARY_NAME)

default: test