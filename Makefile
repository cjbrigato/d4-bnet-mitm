bin:

rebin: bnet build
	go build -o bin/
	GOOS=windows GOARCH=amd64 go build -o bin/

bnet: build

all: bin bnet build
	scripts/rebuild_full
	 
clean: clean-bin clean-generated clean-bnet clean-build

clean-bin:
	rm -rf bin/
clean-bnet:
	rm -rf bnet/
clean-generated:
	rm services/generated.go
clean-pb:
	rm -rf build/pb
clean-proto:
	rm -rf build/protos
clean-buildi:
	rm -rf build/
