build:
	env GOOS=linux GOARCH=amd64 CGOENABLE=0 go build -ldflags="-s -w" -o bin/main api/infra/http/main.go