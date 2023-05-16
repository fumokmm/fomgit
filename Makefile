.PHONY: windows linux darwin

# for Windows
windows:
	GOOS=windows GOARCH=amd64 go build -o fomgit.exe

# for Linux
linux:
	GOOS=linux GOARCH=amd64 go build -o fomgit

# for macOS
darwin:
	GOOS=darwin GOARCH=amd64 go build -o fomgit