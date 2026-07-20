# スタイリングルール

## カラー: セマンティックトークンのみを使う

コンポーネントの SCSS でプリミティブカラー（`v.$mint-500` や `v.$gray-600` など）を直接参照しない。
**必ずセマンティックトークン（`var(--color_*)`）を使う。**

```scss
// NG
.title { color: v.$gray-900; }

// OK
.title { color: var(--color_text-default); }
```

### レイヤー構造

| レイヤー | 場所 | 参照してよい場所 |
|---|---|---|
| プリミティブ (`$mint-500` 等) | `app/styles/variables/_color.scss` | セマンティックトークンの定義時のみ |
| セマンティック (`--color_*`) | `app/styles/foundation/_base.scss` の `:root` | すべてのコンポーネント |

### 欲しい色に合うトークンがないとき

プリミティブに逃げず、`_base.scss` の `:root` にセマンティックトークンを追加してから使う。
命名は `--color_<役割>` で、用途がわかる名前にする（見た目の名前 `--color_green` は禁止）。

```scss
// 例
--color_text-muted: #{v.$gray-500};
--color_border-default: #{v.$gray-200};
--color_button-cta: #{v.$green-700};
```

## スペーシング

8pt グリッド相当の `--space_base`（4px）を基準にする。SCSS では `m.space(n)` を使う。

```scss
padding: #{m.space(2)} #{m.space(5)};  // 8px 20px
```

## タイポグラフィ

フォントサイズ・行間は `m.text-style(<size>)` mixin から取る（`xxs`〜`xxxxl`）。
`font-size` の直書きはしない。

## その他

- スタイルは CSS Modules（`Xxx.module.scss`）。コンポーネントと同階層に置く
- ベースフォントは M PLUS Rounded 1c、ロゴは Paytone One（`__root.tsx` の Google Fonts 読み込みで管理）
