VERSION=1.0.0

GOCMD=go

init:
	$(GOCMD) mod download

gen-proto:
	mkdir -p ./pb/word
	protoc -I ./word --go-grpc_out=./pb --go_out=./pb word/*.proto
