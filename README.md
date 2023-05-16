# fomgit
Fumokmm's Obvious, Minimal git.

# Docker

## Building images and launching containers
```
docker-compose up -d --build
```
## Stop Docker image
```
docker-compose down
or
docker-compose down --rmi all
```

# On Docker container

## Goto /app directory
```
cd /app
```

## Go run
```
go run main.go
```

## Go build


### Example: Generate executable file for Windows.
```
GOOS=windows GOARCH=amd64 go build -o fomgit.exe
```
### Example: Generate executable file for Linux.
```
GOOS=linux GOARCH=amd64 go build -o fomgit
```
### Example: Generate executable file for macOS.
```
GOOS=darwin GOARCH=amd64 go build -o fomgit
```

## Build with Make

### Example: Generate executable file for Windows.
```
make windows
```
### Example: Generate executable file for Linux.
```
make linux
```
### Example: Generate executable file for macOS.
```
make darwin
```
