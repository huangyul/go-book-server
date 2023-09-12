.PHONY:docker
docker:
	@rm webook || true
	@go env -w CGO_ENABLED=0
	@go env -w GOOS=linux
	@go env -w GOARCH=amd64
	@go build -o webook .
	@docker rmi -f jojo/webook:v0.0.1
	@docker build -t jojo/webook:v0.0.1 .
