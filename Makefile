all: prepare fetch-dep bld

prepare: 
	echo "run all"
	mkdir -p ./build/_workspace


bld:
	echo "start build"
	GOPATH=`pwd`/build/_workspace go build -v -o ./nodekeytoaddress ./main.go
	@echo "Done building."


fetch-dep:
	GOPATH=`pwd`/build/_workspace go get -u github.com/ethereum/go-ethereum
	GOPATH=`pwd`/build/_workspace go get -u golang.org/x/crypto/ssh/terminal


clean:
	rm nodekeytoaddress