SHELL := /bin/bash

.PHONY: help schema-export proxy db-password db-connect migrate-status migrate-apply migrate-diff prd-tunnel prd-tunnel-stop prd-db-password prd-db-connect prd-migrate-status prd-migrate-apply local-migrate-new local-migrate-apply local-migrate-status local-migrate-hash local-migrate-reset

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Cloud SQL 接続設定
ENV             ?= dev
GCP_PROJECT     := mahking-$(ENV)
INSTANCE        := mahking-$(ENV)-db
REGION          := asia-northeast1
PROXY_PORT      := 15432
CONNECTION_NAME := $(GCP_PROJECT):$(REGION):$(INSTANCE)

# Schema export commands (single schema: mahking_local)
INITDB_DIR ?= ../../go/postgres-initdb.d
POSTGRES_CONTAINER ?= go-postgres-db-1
POSTGRES_DB ?= postgres
SCHEMA_NAME ?= mahking_local
SCHEMA_FILE ?= 01_schema.sql

schema-export: ## 現在のDBから mahking_local スキーマを抽出して go/postgres-initdb.d へ保存
	@mkdir -p $(INITDB_DIR)
	@echo "Exporting schema '$(SCHEMA_NAME)' from $(POSTGRES_CONTAINER)";
	docker exec $(POSTGRES_CONTAINER) pg_dump -U postgres -d $(POSTGRES_DB) -n $(SCHEMA_NAME) -s > $(INITDB_DIR)/$(SCHEMA_FILE)
	perl -ni -e 'print unless /^SECURITY LABEL FOR anon |^\\(un)?restrict /' $(INITDB_DIR)/$(SCHEMA_FILE)
	perl -pi -e "s/$(SCHEMA_NAME)\\.//g" $(INITDB_DIR)/$(SCHEMA_FILE)
	awk '/CREATE SCHEMA $(SCHEMA_NAME);/ {print; print "SET search_path TO $(SCHEMA_NAME);"; next} 1' $(INITDB_DIR)/$(SCHEMA_FILE) > $(INITDB_DIR)/.tmp_schema && mv $(INITDB_DIR)/.tmp_schema $(INITDB_DIR)/$(SCHEMA_FILE)
	perl -pi -e 's/CREATE SCHEMA $(SCHEMA_NAME);/CREATE SCHEMA IF NOT EXISTS $(SCHEMA_NAME);/g' $(INITDB_DIR)/$(SCHEMA_FILE)

proxy:          ## Cloud SQL Auth Proxy を起動（別ターミナルで実行）
	cloud-sql-proxy $(CONNECTION_NAME) \
		--address 127.0.0.1 \
		--port $(PROXY_PORT)

db-password:    ## Secret Manager から DB パスワードを取得
	@gcloud secrets versions access latest \
		--secret=mahking-$(ENV)-db-password \
		--project=$(GCP_PROJECT)

db-connect:     ## pgcli で Cloud SQL に接続
	@pgcli "postgres://app_user:$$($(MAKE) -s db-password)@localhost:$(PROXY_PORT)/mahking_$(ENV)?sslmode=disable"

migrate-status: ## マイグレーション状態を確認
	@atlas migrate status --env $(ENV) \
		--var "db_password=$$($(MAKE) -s db-password)"

migrate-status-history: ## マイグレーション状態を確認
	@atlas-migrate-status --url "postgres://app_user:$$($(MAKE) -s db-password)@localhost:15432/mahking_dev?search_path=public&sslmode=disable"

migrate-apply:  ## マイグレーションを適用
	@atlas migrate apply --env $(ENV) \
		--var "db_password=$$($(MAKE) -s db-password)"

migrate-diff:   ## スキーマ差分からマイグレーションファイルを生成
	@atlas migrate diff --env $(ENV) \
		--var "db_password=$$($(MAKE) -s db-password)"

# ─── AWS 本番用（SSMトンネル経由）───

AWS_PROFILE  ?= mahking
AWS_REGION   ?= ap-northeast-1
AWS_TF_DIR   := ../terraform/aws/prd
TUNNEL_PORT  := 15433
PRD_DB_PASSWORD_SSM := /mahking/prd/db/password
AWS_CLI := aws --profile $(AWS_PROFILE) --region $(AWS_REGION)

prd-tunnel:          ## [prd] bastionを起動してAuroraへのSSMトンネルを張る（別ターミナルで実行。Ctrl-Cで終了すると自動でbastionを停止）
	@set -eu; \
	id=$$(terraform -chdir=$(AWS_TF_DIR) output -raw bastion_instance_id); \
	host=$$(terraform -chdir=$(AWS_TF_DIR) output -raw aurora_endpoint); \
	trap 'echo; echo "stopping bastion $$id"; $(AWS_CLI) ec2 stop-instances --instance-ids "$$id" --query "StoppingInstances[0].CurrentState.Name" --output text' EXIT; \
	echo "starting bastion $$id"; \
	$(AWS_CLI) ec2 start-instances --instance-ids "$$id" >/dev/null; \
	$(AWS_CLI) ec2 wait instance-running --instance-ids "$$id"; \
	echo "waiting for SSM agent to come online"; \
	status=""; \
	for _ in $$(seq 1 40); do \
		status=$$($(AWS_CLI) ssm describe-instance-information \
			--filters "Key=InstanceIds,Values=$$id" \
			--query 'InstanceInformationList[0].PingStatus' --output text); \
		if [ "$$status" = "Online" ]; then break; fi; \
		sleep 3; \
	done; \
	if [ "$$status" != "Online" ]; then echo "SSM agent is not Online (last: $$status)" >&2; exit 1; fi; \
	echo "localhost:$(TUNNEL_PORT) -> $$host:5432"; \
	$(AWS_CLI) ssm start-session --target "$$id" \
		--document-name AWS-StartPortForwardingSessionToRemoteHost \
		--parameters host="$$host",portNumber=5432,localPortNumber=$(TUNNEL_PORT)

prd-tunnel-stop:     ## [prd] bastionを停止する（prd-tunnelが異常終了して止まらなかったとき用）
	@id=$$(terraform -chdir=$(AWS_TF_DIR) output -raw bastion_instance_id); \
	$(AWS_CLI) ec2 stop-instances --instance-ids "$$id" \
		--query 'StoppingInstances[0].CurrentState.Name' --output text

prd-db-password:     ## [prd] SSM Parameter Store から DB パスワードを取得
	@$(AWS_CLI) ssm get-parameter --name $(PRD_DB_PASSWORD_SSM) \
		--with-decryption --query Parameter.Value --output text

prd-db-connect:      ## [prd] pgcli で Aurora に接続（prd-tunnel を張った状態で実行）
	@pgcli "postgres://mahking:$$($(MAKE) -s prd-db-password)@localhost:$(TUNNEL_PORT)/mahking?sslmode=require"

prd-migrate-status:  ## [prd] マイグレーション状態を確認（prd-tunnel を張った状態で実行）
	@atlas migrate status --env prd \
		--var "db_password=$$($(MAKE) -s prd-db-password)"

prd-migrate-apply:   ## [prd] マイグレーションを適用（prd-tunnel を張った状態で実行）
	@atlas migrate apply --env prd \
		--var "db_password=$$($(MAKE) -s prd-db-password)"

# ─── ローカル用 Atlas コマンド ───

local-migrate-new:    ## [local] マイグレーションファイルの雛形を作成 (name=xxx)
	@atlas migrate new $(name) --env local

local-migrate-apply:  ## [local] マイグレーションを適用
	@atlas migrate apply --env local

local-migrate-status: ## [local] マイグレーション状態を確認
	@atlas migrate status --env local

local-migrate-hash:   ## [local] チェックサムを再計算（マイグレーション手動編集後に実行）
	@atlas migrate hash --env local

local-migrate-reset:  ## [local] スキーマを空にして migrations を最初から全部適用（データは消える）
	@atlas schema clean --env local --auto-approve
	@atlas migrate apply --env local
