# Mahking

麻雀の対局記録アプリ。TanStack Start (React 19 + Vite) + SCSS Modules + Tailwind CSS 4。

## コマンド

- `npm run dev` — 開発サーバ
- `npm run build` — 本番ビルド（`.output/`）+ 型チェック
- `npm run typecheck` — 型チェックのみ

## ルール（必読）

- [ディレクトリ構成](docs/directory-structure.md) — feature 分割 + フラットなコンポーネント構成（`_` プレフィックス = private）
- [スタイリング](docs/styling.md) — 色はセマンティックトークン（`var(--color_*)`）のみ。プリミティブ直接参照禁止

## 補足

- ルート定義は `app/routes/`（TanStack Start のファイルベースルーティング）。ルートは薄く保ち、実装は feature に置く
- `app/routeTree.gen.ts` は自動生成ファイル。編集・コミットしない
- フォームは TanStack Form + zod（`useTanstackForm` hook 経由）
