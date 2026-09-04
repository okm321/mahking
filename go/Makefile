.PHONY: govalid sqlc vet lint con-local-db

govalid: ## govalid のバリデーター生成
	go tool govalid ./...

sqlc: ## sqlc のクエリコード生成
	sqlc generate

vet: ## 静的解析
	go vet ./...

lint: ## golangci-lint 実行
	go tool golangci-lint run ./...

con-local-db: ## ローカルDBにpgcliで接続
	pgcli postgresql://postgres:password@127.0.0.1:15432/postgres

