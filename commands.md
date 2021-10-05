#### Generate asm code
```
GOOS=linux GOARCH=amd64 go tool compile -S main.go
```
or
```
go build -gcflags -S main.go
```
