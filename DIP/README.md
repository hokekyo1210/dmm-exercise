# Solid原則 - DIP

Golangによる例を示す。

## 基礎事項

- DIP = The Dependency Inversion Principle: 依存関係逆転の原則
- > 抽象に対して依存するべきであり、具象に対して依存してはならない。
    - Golangにおいては・・・
        - 全てのパッケージはインタフェースを持つべきで、具体的な細部を知らなくても機能概要を把握できるようにすべきである。
        - パッケージに依存関係が必要となった場合、パラメータとして依存関係を受け入れるべきである。

> ![クラス図](https://miro.medium.com/max/1206/1*unKTIGRnRszZjFWFJa715A.png)

これは参考文献からの引用だが、すべての具体的な実装がインタフェースのみに依存している。

## 例

先程の引用元の画像で挙げた例をGolangで実装してみた。
UserLikeFetcherInterfaceの具体的な実装として, MySQLとPostgresqlを用意した。
まだ良くわかってない部分多し, Javaでやってみるか。


## 参考文献

- [依存関係逆転の原則の重要性について](https://medium.com/eureka-engineering/go-dependency-inversion-principle-8ffaf7854a55)
- [【Go言語】SOLID原則を5分で理解する](https://qiita.com/shunp/items/646c86bb3cc149f7cff9#ddependency-inversion-principle%E4%BE%9D%E5%AD%98%E9%96%A2%E4%BF%82%E9%80%86%E8%BB%A2%E3%81%AE%E5%8E%9F%E5%89%87)
- [SOLID原則を理解するためにGo言語で書いてみた](https://qiita.com/MAKOTODA/items/976f47ea036e35c7538d)