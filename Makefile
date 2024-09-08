.PHONY: docker
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook .
	@docker rmi -f wangzupeng/webook:v0.1.1
	@docker build -t wangzupeng/webook:v0.1.1 .
