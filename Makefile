
GO=go
SRC=passgen.go

default: linux linux-arm freebsd darwin win64 win32

linux:
	GOOS=linux   GOARCH=amd64 $(GO) build -ldflags="-s -w" -o passgen_linux     $(SRC)

linux-arm:
	GOOS=linux   GOARCH=arm   $(GO) build -ldflags="-s -w" -o passgen_linux_arm $(SRC)

freebsd:
	GOOS=freebsd GOARCH=amd64 $(GO) build -ldflags="-s -w" -o passgen_freebsd   $(SRC)

darwin:
	GOOS=darwin  GOARCH=amd64 $(GO) build -ldflags="-s -w" -o passgen_macosx    $(SRC)

win64:
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags="-s -w" -o passgen_win64.exe $(SRC)

win32:
	GOOS=windows GOARCH=386   $(GO) build -ldflags="-s -w" -o passgen_win32.exe $(SRC)

