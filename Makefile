GO=go
GOTEST=$(GO) test
GOBUILD=$(GO) build
MAIN_NAME=./main/main.go
BINARY_NAME=./chord.exe
CHORD_DIR=./chord

PROTOC=protoc
PROTOC_DIR=$(CHORD_DIR)
PROTOC_NAME=$(PROTOC_DIR)/chord.proto


build: clean
	$(PROTOC) --go_out=. --go-grpc_out=require_unimplemented_servers=false:. $(PROTOC_NAME) 
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_NAME)

test: build
	$(BINARY_NAME) -a localhost -p 8001 --ts 3000 --tff 1000 --tcp 3000 -r 4 -v

test2:
	$(BINARY_NAME) -a localhost -p 8002 --ja localhost --jp 8001 --ts 3000 --tff 1000 --tcp 3000 -r 4 -v

test3:
	$(BINARY_NAME) -a localhost -p 8004 --ja localhost --jp 8002 --ts 3000 --tff 1000 --tcp 3000 -r 4 -v

clean:
	rm -f $(BINARY_NAME)

default: test