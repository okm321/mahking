# ディレクトリ構成ルール

bulletproof-react の feature 分割と、フラットなコンポーネントディレクトリ構成
（参考: https://kaminashi-developer.hatenablog.jp/entry/2026/06/05/flat-component-directory）を組み合わせる。

## 全体構成

```
app/
├── routes/          # TanStack Start のルート定義（薄く保つ）
├── features/        # 機能単位のまとまり（本体はここ）
├── components/      # 機能横断で使う共有UIコンポーネント
├── hooks/           # 機能横断で使う共有hooks
├── constants/       # 機能横断で使う定数
├── type-utils/      # 型ユーティリティ
├── styles/          # デザイントークン・mixin・グローバルスタイル
└── router.tsx
```

### routes/ は薄く保つ

ルートファイルの責務は「パス定義・head・feature の View を繋ぐ」だけ。
ロジックや JSX の実装を routes/ に書かない。

```tsx
// app/routes/new.tsx
export const Route = createFileRoute("/new")({
  head: () => ({ meta: [...] }),
  component: () => <GroupCreateView />,
})
```

## features/ の規約

機能（ユーザーから見た関心事）単位でディレクトリを切る。例: `group-create`, `score-input`, `result-dashboard`。

```
features/group-create/
├── index.ts                 # 公開API。外部はここからのみ import する
├── GroupCreateView.tsx      # public なコンポーネント
├── _GroupRuleForm.tsx       # private（feature 内部でのみ使用）
├── _GroupRuleForm.module.scss
├── _useGroupCreateForm.ts   # private な hook
└── schema.ts                # zod スキーマ
```

- feature 間の依存は原則禁止。共有したくなったら `app/components/` や `app/hooks/` に昇格させる
- feature の外から import してよいのは `index.ts` が export するものだけ
- `index.ts` で private（`_` 付き）ファイルを再 export しない

## フラットなコンポーネント構成

**ネストで集約する代わりに、命名で集約する。**

- `_` プレフィックス = private。そのディレクトリの内部実装であり、外から import 禁止
- プレフィックスなし = public
- `components/`, `hooks/`, `model/` のような責務別サブディレクトリを feature 内に作らない。`_Header.tsx`（コンポーネント）、`_useXxx.ts`（hook）のように命名で判別する

### 大きくなったら「ネストさせず、上に切り出す」

1. private ファイルが肥大化したら、深いネストを作らず同階層または1つ上の階層に `_Xxx/` として切り出す
2. さらに大きくなったら sub-feature に昇格する: `features/group/` → `features/group-create/` + `features/group-rule/`

## components/（共有UI）の規約

複数 feature から使われる汎用コンポーネントのみを置く。1コンポーネント = 1ディレクトリ。

```
components/Button/
├── index.ts
├── Button.tsx
└── Button.module.scss
```

特定の機能の知識（ドメイン知識）を持つものは共有化せず、その feature に置く。

## 現状からの移行メモ

既存コードはまだこの構成に揃っていない。触るときに以下の方針で寄せる:

- `views/GroupCreateView` + `components/GroupCreateForm` + `components/FieldGroup*` + `schema/group*` → `features/group-create/` に集約
- `components/shared/*` → `components/*` に昇格（`shared` 階層は廃止）
- 新規コードは最初からこのルールに従う。無関係な既存コードの一斉リネームはしない
