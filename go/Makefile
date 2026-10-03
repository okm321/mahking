.PHONY: sqlc vet lint con-local-db ecr-login image-build image-push

AWS_PROFILE    ?= mahking
AWS_REGION     ?= ap-northeast-1
AWS_ACCOUNT_ID := $(shell aws sts get-caller-identity --profile $(AWS_PROFILE) --query Account --output text)
ECR_REPO       := $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/mahking
IMAGE_TAG      ?= $(shell git rev-parse --short HEAD)

sqlc: ## sqlc のクエリコード生成
	sqlc generate

vet: ## 静的解析
	go vet ./...

lint: ## golangci-lint 実行
	go tool golangci-lint run ./...

con-local-db: ## ローカルDBにpgcliで接続
	pgcli postgresql://postgres:password@127.0.0.1:15432/postgres

ecr-login: ## ECRにdocker login（12時間有効）
	aws ecr get-login-password --profile $(AWS_PROFILE) --region $(AWS_REGION) | docker login --username AWS --password-stdin $(ECR_REPO)

image-build: ## Fargate用（linux/amd64）にイメージをビルド。タグはコミットSHA
	docker build --platform linux/amd64 -t $(ECR_REPO):$(IMAGE_TAG) .

image-push: image-build ## ビルドしてECRにpush
	docker push $(ECR_REPO):$(IMAGE_TAG)
	@echo "pushed: $(ECR_REPO):$(IMAGE_TAG)"

