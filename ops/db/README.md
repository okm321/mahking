# DB スキーマ管理

Atlas（無料版）でマイグレーションを管理する。
`atlas migrate diff` は関数・トリガー・RLS ポリシーが Pro 専用のため使えない。差分 SQL は手書きする。

## ディレクトリ

- `migrations/` — 正。適用する SQL。ファイル名順に流れる。`atlas.sum` でチェックサム管理
- `schema/` — あるべき形の参照用。`diff` には使わないが、テーブル定義の一覧として維持する
- `../../go/postgres-initdb.d/01_schema.sql` — `migrations/` を全部流した結果のスナップショット。sqlc の入力と、`docker compose up` 初回の初期化に使う

## カラムを足す手順（ローカル）

```
1. schema/*.sql を編集（参照用）
2. make local-migrate-new name=add_xxx      # 雛形を作る
3. migrations/<timestamp>_add_xxx.sql に ALTER TABLE を手書き
4. make local-migrate-hash                   # 手書き後にチェックサムを再計算
5. make local-migrate-reset                  # スキーマを空にして migrations を全部流す（データは消える）
6. make schema-export                        # 01_schema.sql を更新
7. cd ../../go && make sqlc                  # 生成コードを更新
```

ローカル DB は「いつでも捨てて `migrations/` から作り直す」前提。テストデータは残らない。
`01_schema.sql` で作った DB に対して `local-migrate-apply` を直接叩くと、Atlas の記録テーブルが空のため失敗する。その場合も `local-migrate-reset` を使う。
