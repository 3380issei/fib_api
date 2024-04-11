# Fibonacci API

この API は、指定された n 番目のフィボナッチ数を返す機能を提供します。

## 使用言語（フレームワーク）

- Go (Gin)

## 構成

クリーンアーキテクチャを採用し、保守運用性や、変更容易性の高い構成にした。

### 各ファイルの責務

- `main.go`: アプリケーションのエントリーポイントで、依存関係の注入と API サーバーの起動を行う。
- `router.go`: エンドポイントのルーティングを行う。
- `controller.go`: HTTP リクエストを処理してバリデーションを行い、適切なレスポンスを返すコントローラが含まれている。
- `usecase.go`: ビジネスロジック（この場合はフィボナッチ数の計算アルゴリズム）を実装するユースケースが含まれている。
- `_test.go`: 各 package のユニットテストが含まれている。
- `_mock.go`: 各 package のモックが含まれている。

## 機能

| 機能                 | メソッド | パス | クエリ |
| -------------------- | -------- | ---- | ------ |
| フィボナッチ数の取得 | GET      | /fib | n      |

### 成功時

リクエスト

```
https://sample.com/fib?n=99
```

レスポンス

```
{
    "result": "218922995834555169026"
}
```

ステータスコード

```
200
```

### エラーの事例（n が整数ではない場合）

リクエスト

```
https://sample.com/fib?n=abc
```

```
https://sample.com/fib?n=3.14
```

レスポンス

```
{
    "message": "Invalid parameter (not integer)"
}
```

ステータスコード

```
400
```

### エラーの事例（n が自然数ではない場合）

リクエスト

```
https://sample.com/fib?n=-1
```

レスポンス

```
{
    "message": "Invalid parameter (not natural number)"
}
```

ステータスコード

```
400
```

## 工夫した点

- 動的計画法を用いたフィボナッチ数の計算アルゴリズムを採用し、再帰だと O(2^n)かかる計算量を O(n)に抑えた。
- big.Int 型でフィボナッチ数を管理することにより、Go の uint 型のオーバーフロー対策をした。
- HTTP レスポンスの型を定義し、予期しないレスポンスを返さないようにした。
- エラー時のメッセージを定数で定義することにより、開発時のタイポを防止した。
- 適切なモック化と依存性の注入をすることで、完全に独立したユニットテストを記述した（100% coverage）。
