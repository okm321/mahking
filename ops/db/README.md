# DB スキーマ管理

Atlas（無料版）でマイグレーションを管理する。
`atlas migrate diff` は関数・トリガー・RLS ポリシーが Pro 専用のため使えない。差分 SQL は手書きする。

## ディレクトリ

- `schema/` — あるべき形の参照用。`diff` には使わないが、テーブル定義の一覧として維持する
- `migrations/` — 適用する SQL。ファイル名順に流れる。`atlas.sum` でチェックサム管理

## カラムを足す手順（ローカル）

```
1. schema/*.sql を編集（参照用）
2. make local-migrate-new name=add_xxx      # 雛形を作る（atlas.sum も更新される）
3. migrations/<timestamp>_add_xxx.sql に ALTER TABLE を手書き
4. make local-migrate-hash                   # 手書き後にチェックサムを再計算
5. make local-migrate-apply                  # ローカル DB に適用
6. make schema-export                        # go/postgres-initdb.d/01_schema.sql を更新
7. cd ../../go && make sqlc                  # 生成コードを更新
```

## 初回セットアップ（initdb で作った DB に Atlas を紐づける）

`docker compose up` 直後の DB は `01_schema.sql` で作られていて Atlas の管理外。
`01_schema.sql` が対応するバージョンを baseline として指定する。

```
atlas migrate apply --env local --baseline <01_schema.sql を export した時点のバージョン>
```
