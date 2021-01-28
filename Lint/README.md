# Lint: 静的コード解析ツール

## 説明

### どのように実行するか

一般的にはエディタのファイル保存時やCIにおいて実行される。

### どのような理由で行うのか

人間系での確認が難しい細かいコーディングミス(コーディング規約違反やコンパイラで見つけられないエラー)を見つける。

### どのような観点に気をつけて実行するのか

実行がもれないようにCIツール(CircleCIやgithub actions)を活用する。

## golangci-lint

api-devにおいて[CircleCI + golangci-lintによる静的コード解析](https://git.dmm.com/platform-gateway/developer-api/blob/ce181d63f1f76343ef060a0a8dc0bb943ca31269/.circleci/config.yml#L10)を行っている。

### 選定理由

他にも代替案として
- govet：Golang デフォルトの Linter
- errcheck：ちゃんとエラーハンドリングしているかチェックしてくれる
- unused：未使用の定義をチェックしてくれる
- goimports：未使用のimportを消してくれたり、フォーマット修正してくれる
- gosimple：コードをシンプルにしてくれる
などがあるが, 全ての機能を包括しているのがgolangci-lintだった。

### 発生した課題

便利すぎて手放せなくなった。

## 参考文献

- [GolangCI-Lintの設定ファイルを理解する DMM Advent Calendar 2019 9日目](https://yyh-gl.github.io/tech-blog/blog/golangci-lint-custom-settings/)