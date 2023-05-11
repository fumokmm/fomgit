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

### for Windows
```
GOOS=windows GOARCH=amd64 go build -o fomgit.exe
```