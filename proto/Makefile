.PHONY: lint generate breaking

lint: ## proto の lint
	buf lint

generate: lint ## lint + Go / TS コード生成
	buf generate

breaking: ## dev ブランチとの破壊的変更チェック
	buf breaking --against '../.git#branch=dev,subdir=proto'
