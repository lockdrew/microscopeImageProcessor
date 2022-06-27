build: 
	go build -o bin/main main.go

run: 
	go run main.go

compile: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags '-extldflags "-static"' -o bin/linux/microscopeImageProcessor main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build --ldflags '-extldflags "-static"' -o bin/windows/microscopeImageProcessor.exe main.go
