# Solid原則 - ISP

Golangによる例を示す。

## 基礎事項

- ISP = Interface Segregation Principle: インタフェース分離の原則
- 不要なインターフェースに依存することを強制してはいけないという原則

## 例

```
type Animal interface {
    breath()
    fly()
    run()
    swim()
}
```
例えば上のようなインタフェースを継承してHuman型を作ろうと思うと, 明らかに不要なfly()の実装を行わなければならない。
そのため以下のようにインタフェースを分割してやるといいだろう。

```
type HumanInterface interface {
	breath()
	run()
	swim()
}

type BirdInterface interface {
	breath()
	fly()
}
```

このように変更することで, Human型はHumanInterfaceを実装することで余計なfly()の実装をせずに済む。

## 他の定義とか

- > クライアントごとに特化したインタフェースを用意する。インタフェースを太くしないということが重要。
- > 簡単に言えば「インターフェースはシンプルにしなさい」と言うことです。

## 参考文献

- [SOLID原則を理解するためにGo言語で書いてみた](https://qiita.com/MAKOTODA/items/976f47ea036e35c7538d)
- [インタフェース分離の原則(ISP)](https://kyog02.hatenablog.com/entry/2019/12/18/001136)
- [オブジェクト指向設計原則とは](https://qiita.com/UWControl/items/98671f53120ae47ff93a#%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%BC%E3%83%95%E3%82%A7%E3%83%BC%E3%82%B9%E5%88%86%E9%9B%A2%E3%81%AE%E6%B3%95%E5%89%87)