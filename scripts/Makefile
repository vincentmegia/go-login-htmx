run-login:
	@templ generate
	@air --build.cmd "go build -o bin/api cmd/login/main.go" --build.bin "./bin/api"
#	@go run cmd/login/main.go

compile:
	@templ generate
	@go build ../cmd/login -ldflags="-X 'github.com/vincentmegia/go-login-htmx/build.Environment=development' -X 'github.com/vincentmegia/go-login-htmx/build.Version=v1.0.0'"
