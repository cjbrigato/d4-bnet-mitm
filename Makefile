runw:
	.\4-bnet-mitm.exe localhost:1119 185.60.112.74:1119

runl:
	 ./bin/d4-bnet-mitm 127.0.0.1:1119 185.60.112.74:1119


bin: rebin

rebin: bnet build
	rsrc.exe -arch amd64 -manifest .\d4-bnet-mitm.exe.manifest -o .\d4-bnet-mitm.syso
	go build
	GOOS=windows GOARCH=amd64 go build

bnet: build

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
