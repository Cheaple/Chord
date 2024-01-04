GO=go
GOTEST=$(GO) test
GOBUILD=$(GO) build
MAIN_NAME=./main/main.go
BINARY_NAME=./chord.exe
CHORD_DIR=./chord
# IP=127.0.0.1
IP=172.23.31.155

PROTOC=protoc
PROTOC_DIR=$(CHORD_DIR)
PROTOC_NAME=$(PROTOC_DIR)/chord.proto

protoc:
	$(PROTOC) --go_out=. --go-grpc_out=require_unimplemented_servers=false:. $(PROTOC_NAME) 

build: clean
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_NAME)

test: build
	$(BINARY_NAME) -a $(IP) -p 8000 --ts 3000 --tff 1000 --tcp 3000 -r 4

test2:
	$(BINARY_NAME) -a $(IP) -p 8002 --ja $(IP) --jp 8000 --ts 3000 --tff 1000 --tcp 3000 -r 4

test3:
	$(BINARY_NAME) -a $(IP) -p 8009 --ja $(IP) --jp 8002 --ts 3000 --tff 1000 --tcp 3000 -r 4

clean:
	rm -f $(BINARY_NAME)

default: test