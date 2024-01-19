build:
	rsrc.exe -arch amd64 -manifest .\d4-bnet-mitm.exe.manifest -o .\d4-bnet-mitm.syso
	go build
	GOOS=windows GOARCH=amd64 go build
	rm .\d4-bnet-mitm.syso
